package main

import (
	"container/list"
)

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