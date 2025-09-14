package main

// kosarajuAlgorithm finds the number of Strongly Connected Components (SCCs) in a directed graph.
// Uses Kosaraju's Algorithm: two DFS passes — first on original graph, second on transpose.
func kosarajuAlgorithm(v int, adjList [][]int) int {
	visited := make([]int, v) // Tracks visited nodes in first DFS
	stack := []int{}          // Stores nodes in order of finishing times (post-order)

	// Step 1: Perform DFS on the original graph to get finishing order
	for i := 0; i < v; i++ {
		if visited[i] == 0 {
			dfs(i, visited, adjList, &stack)
		}
	}

	// Step 2: Build the transpose of the graph (reverse all edges)
	adjListT := make([][]int, v)
	for i := 0; i < v; i++ {
		visited[i] = 0 // Reset visited array for second DFS

		// For every edge u -> v, add edge v -> u in transpose
		for _, neighbor := range adjList[i] {
			adjListT[neighbor] = append(adjListT[neighbor], i)
		}
	}

	ans := 0 // Count of SCCs

	// Step 3: Process nodes in reverse order of finishing times (from stack)
	// Do DFS on the transpose graph — each DFS tree is one SCC
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop from stack

		if visited[node] == 0 {
			// New unvisited node means a new SCC
			ans++
			dfs2(node, visited, adjListT) // Explore entire SCC in transpose graph
		}
	}

	return ans
}

// dfs performs a DFS on the original graph and pushes nodes to stack on backtracking
// This ensures nodes are ordered by decreasing finishing time (post-order)
func dfs(node int, visited []int, adjList [][]int, stack *[]int) {
	visited[node] = 1

	// Visit all unvisited neighbors
	for _, neighbor := range adjList[node] {
		if visited[neighbor] == 0 {
			dfs(neighbor, visited, adjList, stack)
		}
	}

	// After visiting all descendants, push node to stack
	*stack = append(*stack, node)
}

// dfs2 performs DFS on the transpose graph
// In the transpose, a DFS from node 'u' visits exactly all nodes in the same SCC as 'u'
func dfs2(node int, visited []int, adjList [][]int) {
	visited[node] = 1

	// Traverse all reachable nodes in the transpose graph
	for _, neighbor := range adjList[node] {
		if visited[neighbor] == 0 {
			dfs2(neighbor, visited, adjList)
		}
	}
}
