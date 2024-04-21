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
	nodes        []*Node
	adjList      map[string][]string
	visitedCount int
}

func NewGraph() *Graph {
	return &Graph{
		nodes:        []*Node{},
		adjList:      make(map[string][]string),
		visitedCount: 0,
	}
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

// IDDFS implementation
func (g *Graph) IDDFS(startNode string, goalNode string, maxDepth int) []string {
	g.visitedCount = 0

	for depth := 0; depth <= maxDepth; depth++ {
		visited := make(map[string]bool)
		result := g.depthLimitedSearch(startNode, goalNode, depth, visited)
		if len(result) > 0 {
			return result
		}
	}
	return nil
}

// depthLimitedSearch implementation
func (g *Graph) depthLimitedSearch(current string, goal string, depth int, visited map[string]bool) []string {
	if depth == 0 && current == goal {
		return []string{current}
	}
	if depth <= 0 || visited[current] {
		return nil
	}

	visited[current] = true
	for _, neighbor := range g.adjList[current] {
		g.visitedCount++
		if result := g.depthLimitedSearch(neighbor, goal, depth-1, visited); result != nil {
			return append([]string{current}, result...)
		}
	}
	return nil
}

func main() {
	startTitle := "Taylor Swift"
	endTitle := "Fruit of the Loom"

	startURL := convertToURL(startTitle)
	goalURL := convertToURL(endTitle)

	start := time.Now()
	g := NewGraph()

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
				path := g.IDDFS(startURL, goalURL, 3)
				if path != nil {
					fmt.Println("DFS Shortest Path:")
					for _, node := range path {
						fmt.Println(node)
					}
					fmt.Println("Length of Path:", len(path)-1)
					fmt.Println("Number of Articles Visited:", g.visitedCount)
					fmt.Println("DFS Time:", time.Since(start))
					return
				}
			}
			q.PushBack(link)
		}
	}

	fmt.Println("No path found.")
	fmt.Println("Length of Path: 0")
	fmt.Println("Number of Articles Visited:", g.visitedCount)
	fmt.Println("DFS Time:", time.Since(start))
}