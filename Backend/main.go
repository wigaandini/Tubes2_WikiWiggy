package main

import (
	"container/list"
	"fmt"
	"time"
	"strings"
	"bufio"
	"os"
)

func main() {
	var startTitle, goalTitle string
	var choice string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the start title: ")
	startTitle, err := reader.ReadString('\n')
	if err == nil {
		startTitle = strings.TrimSpace(startTitle)
	}
	startURL := convertToURL(startTitle)
	
	fmt.Print("Enter the goal title: ")
	goalTitle, err = reader.ReadString('\n')
	if err == nil {
		goalTitle = strings.TrimSpace(goalTitle)
	}
	goalURL := convertToURL(goalTitle)

	fmt.Print("Enter the searching algorithm: ")
	fmt.Scanln(&choice)

	switch choice {
	case "IDS":
		start := time.Now()
		g := NewGraph()

		visited := make(map[string]bool)
		visited[startURL] = true
		q := list.New()
		q.PushBack(startURL)

		maxDepth := 0

		for q.Len() != 0 {
			currentURL := q.Front().Value.(string)
			q.Remove(q.Front())
			links := linkScraper(currentURL, visited)
			for _, link := range links {
				g.AddEdge(currentURL, link)
				if link == goalURL {
					path := g.IDS(startURL, goalURL, maxDepth)
					for path == nil {
						maxDepth++
						path = g.IDS(startURL, goalURL, maxDepth)
					}
					fmt.Println("IDS Shortest Path:")
					for i, node := range path {
						title := getTitle(node)
						fmt.Printf("%d. %s\n", i+1, strings.ReplaceAll(title, "_", " "))
					}
					fmt.Println("Length of Path:", len(path)-1)
					fmt.Println("Number of Articles Visited:", g.visitedCount)
					fmt.Println("IDS Time:", time.Since(start))
					return
				} 
				q.PushBack(link)
			}
		}

		fmt.Println("No path found.")
		fmt.Println("Length of Path: 0")
		fmt.Println("Number of Articles Visited:", g.visitedCount)
		fmt.Println("IDS Time:", time.Since(start))
	case "BFS":
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
					path := g.BFS(startURL, goalURL)
					if path != nil {
						fmt.Println("BFS Shortest Path:")
						for i, node := range path {
							title := getTitle(node)
							fmt.Printf("%d. %s\n", i+1, strings.ReplaceAll(title, "_", " "))
						}
						fmt.Println("Length of Path:", len(path)-1)
						fmt.Println("Number of Articles Visited:", g.visitedCount) // Output visited count
						fmt.Println("BFS Time:", time.Since(start))
						return
					}
				}
				q.PushBack(link)
			}
		}

		fmt.Println("No path found.")
		fmt.Println("Length of Path: 0")
		fmt.Println("Number of Articles Visited:", g.visitedCount)
		fmt.Println("BFS Time:", time.Since(start))
	default:
		fmt.Println("Wrong method")
	}
}