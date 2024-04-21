package main

import (
	"container/list"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Node struct {
	URL   string
	Depth int
	Nodes []*Node
}

func main() {
	startTitle := "Taylor Swift"
	endTitle := "Fruit of the Loom"

	startURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(startTitle, " ", "_"))
	endURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", strings.ReplaceAll(endTitle, " ", "_"))

	// title = strings.Replace(title, " ", "_", -1)
	maxDepth := 0

	rootNode := &Node{URL: startURL}

	bfsStart := time.Now()
	foundNode, visitedCount := bfs(rootNode, endURL)
	Found := false
	if foundNode != nil {
		Found = false
	}
	for !Found {
		maxDepth += 1
		buildTreeFromLinks(rootNode, 0, maxDepth)
		foundNode, visitedCount = bfs(rootNode, endURL)
		if foundNode != nil {
			Found = true
		}
	}
	if foundNode != nil {
		fmt.Printf("BFS Path from %s to %s:\n", startURL, endURL)
		for i, node := range foundNode {
			fmt.Printf("%d. %s\n", i+1, node.URL)
		}
	} else {
		fmt.Println("End page not found within the specified depth.")
	}

	fmt.Println("BFS Time: ", time.Since(bfsStart))
	fmt.Println("Visited Links Count: ", visitedCount)
}

func bfs(root *Node, endURL string) ([]*Node, int) {
	q := list.New()
	q.PushBack([]*Node{root})

	visitedCount := 0

	for q.Len() > 0 {
		path := q.Remove(q.Front()).([]*Node)
		node := path[len(path)-1]

		if node.URL == endURL {
			return path, visitedCount
		}

		for _, child := range node.Nodes {
			newPath := append(path[:len(path):len(path)], child) // Copy the path slice
			q.PushBack(newPath)
			visitedCount++
		}
	}

	return nil, visitedCount
}

func linkScraper(url string) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var uniqueLinks []string
	visited := make(map[string]bool)

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

func buildTreeFromLinks(node *Node, depth, maxDepth int) {
	if depth >= maxDepth {
		return
	}

	links := linkScraper(node.URL)
	for _, link := range links {
		childNode := &Node{URL: link, Depth: depth + 1}
		node.Nodes = append(node.Nodes, childNode)
		buildTreeFromLinks(childNode, depth+1, maxDepth)
	}
}