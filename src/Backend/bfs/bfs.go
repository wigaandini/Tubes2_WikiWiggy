package main

import (
	"container/list"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// Graph structure
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
			g.visitedCount++
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = currentNode
				q.PushBack(neighbor)
			}
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
		"/wiki/Main_Page",
		"/wiki/Main_Page:",
		"/wiki/Draft:",
		"/wiki/Module:",
		"/wiki/MediaWiki:",
		"/wiki/Index:",
		"/wiki/Education_Program:",
		"/wiki/TimedText:",
		"/wiki/Gadget:",
		"/wiki/Gadget_Definition:",
		"/wiki/Book:",
		"/wiki/AFD:",
		"/wiki/Namespace:",
		"/wiki/Transwiki:",
		"/wiki/Course:",
		"/wiki/Thread:",
		"/wiki/Summary:",
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(link, prefix) {
			return false
		}
	}
	return strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":")
}

var linkCache map[string][]string

func initLinkCache() {
	linkCache = make(map[string][]string)
	file, err := os.Open("cached-bfs.csv")
	if err != nil {
		log.Println("No existing cache file.")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read cache file: ", err)
	}

	for _, record := range records {
		if len(record) >= 2 {
			url := record[0]
			links := record[1:]
			linkCache[url] = links
		}
	}
}

func saveLinkCache() {
	file, err := os.Create("cached-bfs.csv")
	if err != nil {
		log.Fatal("Failed to create cache file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for url, links := range linkCache {
		record := append([]string{url}, links...)
		if err := writer.Write(record); err != nil {
			log.Fatal("Failed to write to cache file: ", err)
		}
	}
}

func linkScraper(url string, visited map[string]bool) []string {
	if links, ok := linkCache[url]; ok {
		return links
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var uniqueLinks []string

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		style, exists := item.Attr("style")
		if exists && (strings.Contains(style, "display: none") || strings.Contains(style, "visibility: hidden")) {
			return
		}

		if _, hiddenExists := item.Attr("hidden"); hiddenExists {
			return
		}

		link, exists := item.Attr("href")
		if !exists || !isValidArticleLink(link) || visited[link] {
			return
		}

		visited[link] = true
		uniqueLinks = append(uniqueLinks, "https://en.wikipedia.org"+link)
	})

	linkCache[url] = uniqueLinks
	return uniqueLinks
}

func getTitle(urlString string) string {
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

func main() {
	initLinkCache()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		startTitle := c.Query("startTitle")
		goalTitle := c.Query("goalTitle")

		if startTitle == "" || goalTitle == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Start title and Goal title are required"})
			return
		}

		start := time.Now()
		g := NewGraph()
		visited := make(map[string]bool)

		startURL := convertToURL(startTitle)
		goalURL := convertToURL(goalTitle)
		visited[startURL] = true
		q := list.New()
		q.PushBack(startURL)

		var mutex sync.Mutex
		pathFound := make(chan []string)

		go func() {
			defer close(pathFound)
			var wg sync.WaitGroup

			for q.Len() != 0 {
				currentURL := q.Front().Value.(string)
				q.Remove(q.Front())

				links := linkScraper(currentURL, visited)
				for _, link := range links {
					wg.Add(1)
					go func(link string) {
						defer wg.Done()
						mutex.Lock()
						defer mutex.Unlock()

						g.AddEdge(currentURL, link)
						if link == goalURL {
							path := g.BFS(startURL, goalURL)
							if path != nil {
								pathFound <- path
								return
							}
						}
						q.PushBack(link)
					}(link)
				}
				wg.Wait()
			}
		}()

		select {
		case path := <-pathFound:
			var pathTitle []string
			for _, node := range path {
				title := getTitle(node)
				pathTitle = append(pathTitle, strings.ReplaceAll(title, "_", " "))
			}
			endTime := time.Since(start).Milliseconds()
			c.JSON(http.StatusOK, gin.H{"paths": pathTitle, "timeTaken": endTime, "visited": g.visitedCount, "length": len(pathTitle) - 1})
		case <-time.After(1000000 * time.Second):
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timed out"})
		}

		saveLinkCache()
	})

	r.Run(":8080")
}