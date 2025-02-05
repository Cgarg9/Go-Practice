package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"os"
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
