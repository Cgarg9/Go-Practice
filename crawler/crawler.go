package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"sync"
	"time"
)

type Crawler struct {
	visited map[string]bool
	mu      sync.Mutex
	wg      sync.WaitGroup
	depth   int
}

func NewCrawler(depth int) *Crawler {
	return &Crawler{
		visited: make(map[string]bool),
		depth:   depth,
	}
}

func (c *Crawler) fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

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

func (c *Crawler) Start(url string) {
	c.wg.Add(1)
	go c.crawl(url, c.depth)
	c.wg.Wait()
}

func main() {
	startURL := "https://example.com"
	depth := 2

	crawler := NewCrawler(depth)
	startTime := time.Now()
	crawler.Start(startURL)
	fmt.Println("Crawling completed in", time.Since(startTime))
}
