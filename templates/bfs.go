package templates

// BFS — Breadth-First Search (Level-Order)
//
// Pattern: Go wide before going deep. Uses queue.
// Time: O(V + E), Space: O(V)

// BFS performs level-order BFS on an adjacency list graph.
// Returns the visit order.
func BFS(graph map[int][]int, start int) []int {
	visited := map[int]bool{start: true}
	queue := []int{start}
	var result []int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

// BFSLevelOrder performs BFS and groups nodes by level.
// Useful for: level-order traversal, shortest path in unweighted graph.
func BFSLevelOrder(graph map[int][]int, start int) [][]int {
	visited := map[int]bool{start: true}
	queue := []int{start}
	var levels [][]int

	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node)

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		levels = append(levels, currentLevel)
	}
	return levels
}

// BFSShortestPath finds the shortest path (by edge count) from start to end.
// Returns the path as a slice of nodes, or nil if no path exists.
func BFSShortestPath(graph map[int][]int, start, end int) []int {
	if start == end {
		return []int{start}
	}

	visited := map[int]bool{start: true}
	parent := map[int]int{start: -1}
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = node
				queue = append(queue, neighbor)

				if neighbor == end {
					// Reconstruct path
					var path []int
					for cur := end; cur != -1; cur = parent[cur] {
						path = append([]int{cur}, path...)
					}
					return path
				}
			}
		}
	}
	return nil // no path
}
