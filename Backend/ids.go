package main

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