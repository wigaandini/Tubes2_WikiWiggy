package main

import (
	"container/list"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

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

// IDS implementation
func (g *Graph) IDS(startNode string, goalNode string, maxDepth int) []string {
	g.visitedCount = 0

	for depth := 0; depth <= maxDepth; depth++ {
		visited := make(map[string]bool)
		result := g.DLS(startNode, goalNode, depth, visited)
		if len(result) > 0 {
			return result
		}
	}
	return nil
}

// DLS implementation
func (g *Graph) DLS(current string, goal string, depth int, visited map[string]bool) []string {
	if depth == 0 && current == goal {
		return []string{current}
	}
	if depth <= 0 || visited[current] {
		return nil
	}

	visited[current] = true
	for _, neighbor := range g.adjList[current] {
		g.visitedCount++
		if result := g.DLS(neighbor, goal, depth-1, visited); result != nil {
			return append([]string{current}, result...)
		}
	}
	return nil
}

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

func getTitle(urlString string) (string) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return ""
	}

	pathParts := strings.Split(parsedURL.Path, "/")
	lastPart := pathParts[len(pathParts)-1]

	title, err := url.PathUnescape(lastPart)
	if err != nil {
		return ""
	}

	return title
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Next()
	}
}

func main () {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		startTitle := c.Query("startTitle")
		goalTitle := c.Query("goalTitle")

		if startTitle == "" || goalTitle == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Start title and Goal title is required"})
			return
		}

		start := time.Now()
		g := NewGraph()

		startURL := convertToURL(startTitle)
		goalURL := convertToURL(goalTitle)

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
					var pathTitle []string
					for _, node := range path {
						title := getTitle(node)
						pathTitle = append(pathTitle, strings.ReplaceAll(title, "_", " "))
					}
					endTime := time.Since(start).Milliseconds()
					c.JSON(http.StatusOK, gin.H{"paths": pathTitle, "timeTaken": endTime, "visited": g.visitedCount, "length": len(pathTitle) - 1})
					return
				} 
				q.PushBack(link)
			}
		}
		c.JSON(http.StatusOK, gin.H{"paths": "", "timeTaken": time.Since(start).Milliseconds(), "visited": g.visitedCount, "length": 0})
	})
	r.Run(":8080") 
}