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
	startURL := "https://en.wikipedia.org/wiki/Artificial_intelligence"
	endURL := "https://en.wikipedia.org/wiki/Machine_learning"

	// title = strings.Replace(title, " ", "_", -1)
	maxDepth := 1

	rootNode := &Node{URL: startURL}
	buildTreeFromLinks(rootNode, 0, maxDepth)

	bfsStart := time.Now()
	foundNode := bfs(rootNode, endURL)
	Found := false
	if foundNode != nil {
		Found = false
	}
	for !Found {
		maxDepth += 1
		buildTreeFromLinks(rootNode, 0, maxDepth)
		foundNode = bfs(rootNode, endURL)
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

	dfsStart := time.Now()
	path := dfs(rootNode, endURL, 0, maxDepth, []*Node{})
	if len(path) > 0 {
		fmt.Printf("DFS path from %s to %s within depth %d:\n", startURL, endURL, maxDepth)
		for i, node := range path {
			fmt.Printf("%d. %s\n", i+1, node.URL)
		}
	} else {
		fmt.Println("Path not found within the specified depth.")
	}
	fmt.Println("DFS Time: ", time.Since(dfsStart))
}

func bfs(root *Node, endURL string) []*Node {
	q := list.New()
	q.PushBack([]*Node{root})

	for q.Len() > 0 {
		path := q.Remove(q.Front()).([]*Node)
		node := path[len(path)-1]

		if node.URL == endURL {
			return path
		}

		for _, child := range node.Nodes {
			newPath := append(path[:len(path):len(path)], child) // Copy the path slice
			q.PushBack(newPath)
		}
	}

	return nil
}

func dfs(node *Node, endURL string, currentDepth, maxDepth int, currentPath []*Node) []*Node {
	if node.URL == endURL {
		return append(currentPath, node)
	}

	if currentDepth < maxDepth {
		for _, child := range node.Nodes {
			if child.URL == node.URL { // Skip self link
				continue
			}
			newPath := append(currentPath, node)
			foundPath := dfs(child, endURL, currentDepth+1, maxDepth, newPath)
			if len(foundPath) > 0 {
				return foundPath
			}
		}
	}

	if node.URL == endURL && currentDepth >= maxDepth {
		return nil
	}

	return nil
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
		if strings.HasPrefix(link, "/wiki/") && !visited[link] {
			visited[link] = true
			// fmt.Printf("Link: '%s'\n", "https://en.wikipedia.org"+link)
			uniqueLinks = append(uniqueLinks, "https://en.wikipedia.org"+link)
		}
	})

	// for i, link := range uniqueLinks {
	// 	fmt.Printf("Link #%d: '%s'\n", i, link)
	// }

	return uniqueLinks
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

// func printTree(node *Node) {
// 	printNode(node, 0)
// }

// func printNode(node *Node, level int) {
// 	fmt.Printf("%s- %s\n", strings.Repeat("\t", level), node.URL)
// 	for _, child := range node.Nodes {
// 		printNode(child, level+1)
// 	}
// }
