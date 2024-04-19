package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

// Function to fetch links from a Wikipedia article
func fetchLinks(url string) []string {
    var links []string

    // Make HTTP request to fetch the page content
    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    // Find all links in the document
    doc.Find("#content a").Each(func(i int, s *goquery.Selection) {
        href, _ := s.Attr("href")
        if strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, ":") {
            links = append(links, "https://en.wikipedia.org"+href)
        }
    })

    return links
}

// Implementing the IDS algorithm
func iterativeDeepeningSearch(startURL, targetURL string, depth int) []string {
    visited := make(map[string]bool)
    path := []string{startURL}
    return depthLimitedSearch(startURL, targetURL, depth, visited, path)
}

func depthLimitedSearch(node, targetURL string, depth int, visited map[string]bool, path []string) []string {
    if depth == 0 {
        return nil
    }
    links := fetchLinks(node)
    for _, link := range links {
        if link == targetURL {
            return append(path, targetURL)
        }
        if !visited[link] {
            visited[link] = true
            newPath := append(path, link)
            result := depthLimitedSearch(link, targetURL, depth-1, visited, newPath)
            if result != nil {
                return result
            }
        }
    }
    return nil
}

func main() {
    // Example usage
    var startURL, targetURL string
    fmt.Print("Enter the start Wikipedia article URL: ")
    fmt.Scanln(&startURL)
    fmt.Print("Enter the target Wikipedia article URL: ")
    fmt.Scanln(&targetURL)

    // Implement IDS algorithm to find shortest path
    path := iterativeDeepeningSearch(startURL, targetURL, 10)
    if path != nil {
        fmt.Println("Shortest path:")
        for _, link := range path {
            fmt.Println(link)
        }
    } else {
        fmt.Println("No path found within the depth limit.")
    }
}
