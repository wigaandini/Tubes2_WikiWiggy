package main

import (
	"container/list"
	"fmt"
	"time"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)


// Link Scraper
func convertToURL(title string) string {
	return fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(title, " ", "_"))
}

func isValidArticleLink(link string) bool {
	prefixes := []string{
		"/wiki/Special:",
		"/wiki/Talk:",
		"/wiki/User:",
		"/wiki/Portal:",
		"/wiki/Wikipedia:",
		"/wiki/File:",
		"/wiki/Category:",
		"/wiki/Help:",
		"/wiki/Template:",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(link, prefix) {
			return false
		}
	}
	return strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":")
}

func linkScraper(url string, visited map[string]bool) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var uniqueLinks []string

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		if isValidArticleLink(link) && !visited[link] {
			visited[link] = true
			uniqueLinks = append(uniqueLinks, "https://en.wikipedia.org"+link)
		}
	})

	return uniqueLinks
}


// Graph
type Graph struct {
	nodes   []*Node
	adjList map[string][]string
}

type Node struct {
	val string
}

func (g *Graph) AddNode(value string) *Node {
	node := &Node{val: value}
	g.nodes = append(g.nodes, node)
	return node
}

func (g *Graph) AddEdge(node1, node2 string) {
	g.adjList[node1] = append(g.adjList[node1], node2)
	g.adjList[node2] = append(g.adjList[node2], node1)
}


// BFS
func (g *Graph) BFS(start, end string) []string {
	visited := make(map[string]bool)
	parent := make(map[string]string)
	q := list.New()

	visited[start] = true
	q.PushBack(start)

	for q.Len() != 0 {
		currentNode := q.Front().Value.(string)
		q.Remove(q.Front())

		if currentNode == end {
			path := []string{}
			current := end
			for current != "" {
				path = append([]string{current}, path...)
				current = parent[current]
			}
			return path
		}

		for _, neighbor := range g.adjList[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = currentNode
				q.PushBack(neighbor)
			}
		}
	}

	return nil
}


// Main
func main() {
    startTitle := "Taylor Swift"
	endTitle := "Fruit of the Loom"

	startURL := convertToURL(startTitle)
	goalURL := convertToURL(endTitle)

    start := time.Now()
    g := &Graph{adjList: make(map[string][]string)}
    visited := make(map[string]bool)

    visited[startURL] = true
    q := list.New()
    q.PushBack(startURL)

    for q.Len() != 0 {
        currentURL := q.Front().Value.(string)
        q.Remove(q.Front())

        links := linkScraper(currentURL, visited)
        for _, link := range links {
            g.AddEdge(currentURL, link)
            if link == goalURL {
                path := g.BFS(startURL, goalURL)
                if path != nil {
                    fmt.Println("BFS Shortest Path:")
                    for _, node := range path {
                        fmt.Println(node)
                    }
                    fmt.Println("BFS Time:", time.Since(start))
                    return
                }
            }
            q.PushBack(link)
        }
    }

    fmt.Println("No path found.")
    fmt.Println("BFS Time:", time.Since(start))
}