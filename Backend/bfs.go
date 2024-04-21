package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type URLQueue struct {
	links         []string
	visited       map[string]bool
	path          map[string]string // To store the path
	currentLink   string             // To keep track of the current link being visited
	neighborLinks []string
}

func (q *URLQueue) Enqueue(link string, parentLink string) {
	q.links = append(q.links, link)
	q.path[link] = parentLink
}

func (q *URLQueue) Dequeue() string {
	if len(q.links) != 0 {
		link := q.links[0]
		q.links = q.links[1:]
		return link
	}
	return ""
}

func (q *URLQueue) IsEmpty() bool {
	return len(q.links) == 0
}

func validLink(link string) bool {
	invalidPrefixes := []string{"/wiki/Special:", "/wiki/Talk:", "/wiki/User:", "/wiki/Portal:", "/wiki/Wikipedia:", "/wiki/File:", "/wiki/Category:", "/wiki/Help:"}
	for _, prefix := range invalidPrefixes {
		if strings.HasPrefix(link, prefix) {
			return false
		}
	}
	return strings.HasPrefix(link, "/wiki/")
}

func BFS(startURL string, targetURL string) ([]string, time.Duration) {
	urlQueue := URLQueue{
		visited: make(map[string]bool),
		path:    make(map[string]string),
	}

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	startTime := time.Now()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		urlQueue.visited[r.URL.String()] = true
		urlQueue.currentLink = r.URL.String() // Update current link being visited
		urlQueue.Enqueue(r.URL.String(), urlQueue.path[urlQueue.currentLink])
	})

	c.OnHTML("div#mw-content-text "+"a[href]", func(e *colly.HTMLElement) {
		neighborLink := e.Attr("href")
		if validLink(neighborLink) && !urlQueue.visited[neighborLink] {
			urlQueue.neighborLinks = append(urlQueue.neighborLinks, e.Request.AbsoluteURL(neighborLink))
		}
	})

	c.Visit(startURL)

	for !urlQueue.IsEmpty() {
		currentURL := urlQueue.Dequeue()

		if currentURL == targetURL {
			// Reconstruct the path
			path := []string{currentURL}
			for parentURL := urlQueue.path[currentURL]; parentURL != ""; parentURL = urlQueue.path[parentURL] {
				path = append([]string{parentURL}, path...)
			}
			duration := time.Since(startTime)
			return path, duration
		}

		for _, neighborLink := range urlQueue.neighborLinks {
			if !urlQueue.visited[neighborLink] && !contains(urlQueue.links, neighborLink) {
				urlQueue.currentLink = neighborLink // Update current link being visited
				c.Visit(neighborLink)
			}
		}
	}

	return nil, 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	path, duration := BFS("https://en.wikipedia.org/wiki/Artificial_intelligence", "https://en.wikipedia.org/wiki/Machine_learning")
	if path != nil {
		fmt.Println("Shortest path:", path)
		fmt.Println("Time taken:", duration)
	} else {
		fmt.Println("No path found")
	}
}
