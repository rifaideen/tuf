package main

// kosarajuAlgorithm returns all Strongly Connected Components (SCCs) in the graph.
// Each SCC is represented as a list of nodes.
func kosarajuAlgorithm(v int, adjList [][]int) [][]int {
	visited := make([]int, v)
	stack := []int{}

	// Step 1: First DFS on original graph → fill stack by finishing time
	for i := range v {
		if visited[i] == 0 {
			dfs(i, visited, adjList, &stack)
		}
	}

	// Step 2: Build transpose graph (reverse all edges)
	adjListT := make([][]int, v)

	for i := range v {
		visited[i] = 0 // Reset visited for second pass

		for _, neighbor := range adjList[i] {
			adjListT[neighbor] = append(adjListT[neighbor], i)
		}
	}

	var sccs [][]int // To store all SCCs

	// Step 3: Process nodes in reverse finishing order
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] == 0 {
			var component []int                       // Collect nodes in this SCC
			dfs2(node, visited, adjListT, &component) // Fill component
			sccs = append(sccs, component)
		}
	}

	return sccs
}

// dfs: Standard DFS to fill stack in post-order (finishing times)
func dfs(node int, visited []int, adjList [][]int, stack *[]int) {
	visited[node] = 1

	for _, neighbor := range adjList[node] {
		if visited[neighbor] == 0 {
			dfs(neighbor, visited, adjList, stack)
		}
	}

	*stack = append(*stack, node) // Push when backtracking
}

// dfs2: DFS on transpose graph that collects all nodes in current SCC
func dfs2(node int, visited []int, adjList [][]int, component *[]int) {
	visited[node] = 1
	*component = append(*component, node) // Add node to current SCC

	for _, neighbor := range adjList[node] {
		if visited[neighbor] == 0 {
			dfs2(neighbor, visited, adjList, component)
		}
	}
}
