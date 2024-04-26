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

	var wg sync.WaitGroup
	var mutex sync.Mutex


	run := func(current string, goal string, depth int, visited map[string]bool) []string {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		return g.DLS(current, goal, depth, visited)
	}

	for depth := 0; depth <= maxDepth; depth++ {
		visited := make(map[string]bool)
		wg.Add(1)
		result := run(startNode, goalNode, depth, visited)
		wg.Wait()
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
	file, err := os.Open("cached-ids.csv")
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
	file, err := os.Create("cached-ids.csv")
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
	initLinkCache()
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

		if startTitle == goalTitle {
			var paths []string
			paths = append(paths, startTitle)
			c.JSON(http.StatusOK, gin.H{"paths": paths, "timeTaken": time.Since(start).Milliseconds(), "visited": 0, "length": 0})
			return
		}

		startURL := convertToURL(startTitle)
		goalURL := convertToURL(goalTitle)

		visited := make(map[string]bool)
		visited[startURL] = true
		q := list.New()
		q.PushBack(startURL)

		maxDepth := 0

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
							path := g.IDS(startURL, goalURL, maxDepth)
							for path == nil {
								maxDepth++
								path = g.IDS(startURL, goalURL, maxDepth)
								if path != nil {
									pathFound <- path
									return
								}
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
	r.Run(":8081") 
}