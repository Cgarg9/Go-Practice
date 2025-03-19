# Go Web Crawler

## Overview
This project is a basic concurrent web crawler implemented in Go. It fetches and parses web pages, extracts links, and recursively visits them up to a specified depth using Go's concurrency features.

## Features
✅ Efficient crawling with Goroutines and Channels  
✅ Depth control to limit recursive crawling  
✅ Avoids revisiting the same URL  
✅ Parses and extracts links from HTML pages  
✅ Measures execution time for performance tracking  

1. Clone the repository:
   ```bash
   git clone https://github.com/Cgarg9/Go-Practice.git
   cd GO_PRACTICE/crawler
   go mod init crawler
    ```
2. Install dependencies :
    ```bash
    go mod tidy
    ```
3. Run the application:
    ```
    go run crawler.go
    ```

You can customize the `startURL` and `depth` in the `main.go` file:
```go
startURL := "https://example.com"
depth := 2
```

## Code Structure & Explanation

### **Crawler Struct**
Manages visited URLs and depth control.
```go
type Crawler struct {
	visited map[string]bool
	mu      sync.Mutex
	wg      sync.WaitGroup
	depth   int
}
```
- `visited`: Keeps track of URLs that have been crawled.
- `mu`: Ensures safe concurrent access to shared data.
- `wg`: Synchronizes goroutines.
- `depth`: Defines the crawling depth.

### **NewCrawler Function**
Creates a new instance of the crawler.
```go
func NewCrawler(depth int) *Crawler {
	return &Crawler{
		visited: make(map[string]bool),
		depth:   depth,
	}
}
```

### **fetch Function**
Fetches the HTML content of a given URL.
```go
func (c *Crawler) fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
```
- Uses `http.Get()` to retrieve the webpage.
- Returns the HTTP response or an error.

### **parse Function**
Parses the response and extracts links.
```go
func (c *Crawler) parse(resp *http.Response) []string {
	var links []string
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return links
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}
```
- Parses the HTML document.
- Extracts anchor (`<a>`) tag links.

### **crawl Function**
Recursively fetches and parses links up to a specified depth.
```go
func (c *Crawler) crawl(url string, depth int) {
	defer c.wg.Done()
	if depth <= 0 {
		return
	}
	c.mu.Lock()
	if c.visited[url] {
		c.mu.Unlock()
		return
	}
	c.visited[url] = true
	c.mu.Unlock()

	fmt.Println("Fetching:", url)
	resp, err := c.fetch(url)
	if err != nil {
		fmt.Println("Error fetching:", err)
		return
	}
	defer resp.Body.Close()

	links := c.parse(resp)
	for _, link := range links {
		c.wg.Add(1)
		go c.crawl(link, depth-1)
	}
}
```
- Checks if the URL has been visited.
- Fetches and parses the page.
- Recursively follows links.

### **Start Function**
Initiates the crawling process.
```go
func (c *Crawler) Start(url string) {
	c.wg.Add(1)
	go c.crawl(url, c.depth)
	c.wg.Wait()
}
```
- Starts crawling from the given URL.
- Uses `WaitGroup` to handle concurrent operations.

## Dependencies
This project uses:
- `golang.org/x/net/html` for HTML parsing

Install dependencies using:
```sh
go get golang.org/x/net/html
```

## Contribution
Feel free to submit issues or pull requests to improve the crawler.

## License
This project is open-source and available under the MIT License.

