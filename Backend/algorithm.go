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

 maxDepth := 1

 rootNode := &Node{URL: startURL}
 buildTreeFromLinks(rootNode, 0, maxDepth)
 
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
 
 bfsStart := time.Now()
 foundNode := bfsWithIncreasingDepth(rootNode, endURL, maxDepth)
 if foundNode != nil {
  fmt.Printf("BFS Path from %s to %s:\n", startURL, endURL)
  for i, node := range foundNode {
   fmt.Printf("%d. %s\n", i+1, node.URL)
  }
 } else {
  fmt.Println("End page not found within the specified depth.")
 }
 fmt.Println("BFS Time: ", time.Since(bfsStart))
}

func bfsWithIncreasingDepth(root *Node, endURL string, maxDepth int) []*Node {
 q := list.New()
 q.PushBack([]*Node{root})

 visited := make(map[string]bool)
 visited[root.URL] = true

 for q.Len() > 0 {
  path := q.Remove(q.Front()).([]*Node)
  node := path[len(path)-1]

  if node.URL == endURL {
   return path
  }

  // Continue building the tree beyond the initial maxDepth
  if len(path) <= maxDepth {
   links := linkScraper(node.URL, visited)
   for _, link := range links {
    childNode := &Node{URL: link, Depth: node.Depth + 1}
    node.Nodes = append(node.Nodes, childNode)
    newPath := append(path, childNode)
    q.PushBack(newPath)
   }
  }
 }

 return nil
}

func dfs(node *Node, endURL string, currentDepth, maxDepth int, currentPath []*Node) []*Node {
 if node.URL == endURL {
  return append(currentPath, node)
 }

 if currentDepth < maxDepth {
  shortestPath := make([]*Node, 0)

  for _, child := range node.Nodes {
   if child.URL == node.URL { // Skip self link
    continue
   }

   newPath := append(currentPath, node)
   foundPath := dfs(child, endURL, currentDepth+1, maxDepth, newPath)

   if len(foundPath) > 0 && (len(shortestPath) == 0 || len(foundPath) < len(shortestPath)) {
    shortestPath = foundPath
   }
  }

  return shortestPath
 }

 return nil
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
  if strings.HasPrefix(link, "/wiki/") && !visited[link] {
   visited[link] = true
   uniqueLinks = append(uniqueLinks, "https://en.wikipedia.org"+link)
  }
 })

 return uniqueLinks
}

func buildTreeFromLinks(node *Node, depth, maxDepth int) {
 if depth >= maxDepth {
  return
 }

 visited := make(map[string]bool)
 visited[node.URL] = true
 links := linkScraper(node.URL, visited)
 for _, link := range links {
  childNode := &Node{URL: link, Depth: depth + 1}
  node.Nodes = append(node.Nodes, childNode)
  buildTreeFromLinks(childNode, depth+1, maxDepth)
 }
}