package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

// Function to scrape links from a Wikipedia article
func getLinksFromWikipedia(url string) []string {
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    links := []string{}
    if resp.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("a[href^='/wiki/']").Each(func(i int, s *goquery.Selection) {
        link, _ := s.Attr("href")
        if !strings.Contains(link, ":") { // Exclude links with colons (administrative links)
            links = append(links, link)
        }
    })

    return links
}

// Function to find the shortest path using Iterative Deepening Search (IDS)
func iterativeDeepeningSearch(startURL, goalURL string, maxDepth int) []string {
    visited := make(map[string]bool)
    queue := make([][]string, 0)
    queue = append(queue, []string{startURL})

    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]

        currentURL := path[len(path)-1]
        if currentURL == goalURL {
            return path
        }

        if len(path) <= maxDepth {
            links := getLinksFromWikipedia(currentURL)
            for _, link := range links {
                if !visited[link] {
                    newPath := append([]string(nil), path...)
                    newPath = append(newPath, link)
                    queue = append(queue, newPath)
                    visited[link] = true
                }
            }
        }
    }

    return nil
}

func main() {
    var startURL, goalURL string
    fmt.Print("Enter the start Wikipedia article URL: ")
    fmt.Scanln(&startURL)
    fmt.Print("Enter the target Wikipedia article URL: ")
    fmt.Scanln(&goalURL)

    maxDepth := 10 // You can adjust this based on your preference

    path := iterativeDeepeningSearch(startURL, goalURL, maxDepth)

    if path != nil {
        fmt.Println("Shortest path found:")
        for _, link := range path {
            fmt.Println("https://en.wikipedia.org" + link)
        }
    } else {
        fmt.Println("No path found within the specified depth limit.")
    }
}
