package templates

// DFS — Depth-First Search (Recursive + Iterative)
//
// Pattern: Go deep before going wide. Uses stack (explicit or call stack).
// Time: O(V + E), Space: O(V)

// DFSRecursive performs recursive DFS on an adjacency list graph.
func DFSRecursive(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	var result []int
	dfHelper(graph, start, visited, &result)
	return result
}

func dfHelper(graph map[int][]int, node int, visited map[int]bool, result *[]int) {
	if visited[node] {
		return
	}
	visited[node] = true
	*result = append(*result, node)

	for _, neighbor := range graph[node] {
		dfHelper(graph, neighbor, visited, result)
	}
}

// DFSIterative performs iterative DFS using an explicit stack.
// Useful when recursion depth could cause stack overflow.
func DFSIterative(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	stack := []int{start}
	var result []int

	for len(stack) > 0 {
		// Pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}
		visited[node] = true
		result = append(result, node)

		// Push neighbors in reverse for consistent ordering
		neighbors := graph[node]
		for i := len(neighbors) - 1; i >= 0; i-- {
			if !visited[neighbors[i]] {
				stack = append(stack, neighbors[i])
			}
		}
	}
	return result
}

// HasCycleDFS detects a cycle in a directed graph using DFS coloring.
// Colors: 0=white (unvisited), 1=gray (in progress), 2=black (done)
func HasCycleDFS(graph map[int][]int, numNodes int) bool {
	color := make(map[int]int) // default 0 = white

	for node := 0; node < numNodes; node++ {
		if color[node] == 0 {
			if dfsCycleHelper(graph, node, color) {
				return true
			}
		}
	}
	return false
}

func dfsCycleHelper(graph map[int][]int, node int, color map[int]int) bool {
	color[node] = 1 // gray — in progress

	for _, neighbor := range graph[node] {
		if color[neighbor] == 1 {
			return true // back edge → cycle
		}
		if color[neighbor] == 0 {
			if dfsCycleHelper(graph, neighbor, color) {
				return true
			}
		}
	}

	color[node] = 2 // black — done
	return false
}
