package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var store = make(map[string]string)
var mu sync.Mutex

const storeFile = "store.json"

var randGen = rand.New(rand.NewSource(time.Now().UnixNano())) // Local random generator

func init() {
	loadStore()
}

// isValidURL checks if the input string is a valid URL with http or https scheme.
func isValidURL(input string) bool {
	parsedURL, err := url.Parse(input)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	// Allow only "http" and "https" schemes
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	return true
}

// checkHostResolvable checks if the host in the URL is resolvable using net.LookupHost().
func checkHostResolvable(input string) error {
	parsedURL, err := url.Parse(input)
	if err != nil {
		return err
	}

	// Try resolving the host
	_, err = net.LookupHost(parsedURL.Host)
	if err != nil {
		return fmt.Errorf("could not resolve host: %v", err)
	}

	return nil
}

func generateShortURL() string {
	shortURL := ""
	maxAttempts := 100 // Set a maximum number of attempts
	attempts := 0

	for {
		attempts++
		if attempts > maxAttempts {
			fmt.Println("Max attempts reached, unable to generate unique URL")
			return "" // Fail gracefully if we can't generate a unique URL
		}

		b := make([]rune, 6)
		for i := range b {
			b[i] = letters[randGen.Intn(len(letters))]
		}
		shortURL = string(b)

		fmt.Println("Attempting to generate short URL:", shortURL)

		mu.Lock() // Lock mutex for checking uniqueness
		_, exists := store[shortURL]
		mu.Unlock() // Unlock immediately after checking

		if !exists {
			fmt.Println("Unique short URL generated:", shortURL)
			break
		}
		fmt.Println("Collision detected, retrying...")
	}

	return shortURL
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request at /shorten")

	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error decoding request:", err.Error())
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		fmt.Println("Empty URL received")
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Validate the URL
	if !isValidURL(req.URL) {
		// If the URL is missing the protocol, prepend "http://"
		if !isValidURL("http://" + req.URL) {
			fmt.Println("Invalid URL received:", req.URL)
			http.Error(w, "Invalid URL. Please provide a valid URL with http:// or https://", http.StatusBadRequest)
			return
		}
		req.URL = "http://" + req.URL
	}

	// Check if the host is resolvable
	if err := checkHostResolvable(req.URL); err != nil {
		fmt.Println("Host resolution error:", err)
		http.Error(w, fmt.Sprintf("Could not resolve host: %s", err), http.StatusBadRequest)
		return
	}

	fmt.Println("Shortening URL:", req.URL)

	shortURL := generateShortURL()
	if shortURL == "" {
		fmt.Println("Failed to generate unique short URL")
		http.Error(w, "Failed to generate unique URL", http.StatusInternalServerError)
		return
	}

	fmt.Println("Storing short URL:", shortURL)
	mu.Lock() // Lock mutex while updating the store
	store[shortURL] = req.URL
	mu.Unlock() // Unlock after updating the store

	saveStore() // Save store to file

	fmt.Println("Store after saving:", store) // Debugging the store content

	resp := map[string]string{"short_url": "http://localhost:8080/" + shortURL}
	fmt.Println("Generated short URL:", resp["short_url"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func saveStore() {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.Marshal(store)
	if err != nil {
		fmt.Println("Error marshalling store:", err)
		return
	}

	err = os.WriteFile(storeFile, data, 0644) // Use os.WriteFile instead of ioutil.WriteFile
	if err != nil {
		fmt.Println("Error saving store to file:", err)
	} else {
		fmt.Println("Store data saved successfully")
	}
}

func loadStore() {
	mu.Lock()
	defer mu.Unlock()

	data, err := os.ReadFile(storeFile)
	if err != nil {
		// If there's an error reading the file, initialize an empty store
		if !os.IsNotExist(err) {
			fmt.Println("Error reading store file:", err)
		}
		return
	}

	// If the file is empty, just return and initialize an empty store
	if len(data) == 0 {
		return
	}

	// Attempt to unmarshal the data, handle errors if any
	if err := json.Unmarshal(data, &store); err != nil {
		fmt.Println("Error unmarshalling store:", err)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	longURL, exists := store[r.URL.Path[1:]]
	mu.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	fmt.Println("Server started on port 8080")
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/", redirectHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
