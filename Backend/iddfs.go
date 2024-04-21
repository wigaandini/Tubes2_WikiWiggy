package main

import (
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
	temp := 0
	maxDepth := 4

	rootNode := &Node{URL: startURL}


	dfsStart := time.Now()
	path, visitedCount := iddfs(rootNode, endURL, 0, temp, []*Node{}, 0)
	Found := false
	if path != nil {
		Found = false
	}
	for !Found {
		temp += 1
		if temp <= maxDepth {
			buildTreeFromLinks(rootNode, 0, temp)
			path, visitedCount = iddfs(rootNode, endURL, 0, temp, []*Node{}, 0)
			if path != nil {
				Found = true
			}
		}
	}
	if path != nil {
		fmt.Printf("DFS path from %s to %s:\n", startURL, endURL)
		for i, node := range path {
			fmt.Printf("%d. %s\n", i+1, node.URL)
		}
	} else {
		fmt.Println("Path not found within the specified depth.")
	}
	fmt.Println("DFS Time: ", time.Since(dfsStart))
	fmt.Println("Visited Links Count: ", visitedCount)
}

func iddfs(node *Node, endURL string, currentDepth, maxDepth int, currentPath []*Node, visitedCount int) ([]*Node, int) {
	if node.URL == endURL && currentDepth <= maxDepth{
		return append(currentPath, node), visitedCount
	}

	if currentDepth < maxDepth {
		for _, child := range node.Nodes {
			if child.URL == node.URL { // Skip self link
				continue
			}
			newPath := append(currentPath, node)
			visitedCount++
			foundPath, visitedCount := iddfs(child, endURL, currentDepth+1, maxDepth, newPath, visitedCount)
			if len(foundPath) > 0 {
				return foundPath, visitedCount
			}
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