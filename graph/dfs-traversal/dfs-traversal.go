package main

// represent builds an adjacency list representation of an undirected graph.
// Why? Because graphs are easier to traverse when we store each node's neighbors in a list.
// Input:
//   n = number of nodes (nodes are labeled 1 to n)
//   list = list of edges, each edge [u, v] means u is connected to v
// Output: adjacency list where graph[i] contains all neighbors of node i
func represent(n int, list [][]int) [][]int {
	// Initialize an adjacency list with n+1 slots (1-indexed, so we ignore index 0)
	graph := make([][]int, n+1)

	// For each edge [u, v], add v to u's list and u to v's list (undirected graph)
	for _, info := range list {
		i, j := info[0], info[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}

	return graph
}

// dfsTraversal performs a Depth-First Search starting from a given node.
// It returns the order in which nodes are visited.
// Intuition:
//   DFS goes as deep as possible along one path before backtracking.
//   We use recursion and a visited array to avoid revisiting nodes and getting stuck in cycles.
func dfsTraversal(n, start int, list [][]int) []int {
	ans := []int{} // This will store the DFS visit order

	// Build the graph from the edge list
	graph := represent(n, list)

	// Track visited nodes to prevent revisiting and infinite loops
	visited := make([]int, n+1)

	// Mark the starting node as visited
	visited[start] = 1

	// Define the recursive traversal function
	// Why recursion? DFS naturally matches recursion: go deep, then backtrack.
	var traverse func(node int)
	traverse = func(node int) {
		// Visit the current node: record it in the answer
		ans = append(ans, node)

		// Explore all unvisited neighbors
		// Note: The order depends on how neighbors are stored in the adjacency list.
		// If edges were added in left-to-right order, and we assume that's preserved,
		// then DFS will follow that direction.
		for _, neighbor := range graph[node] {
			if visited[neighbor] == 0 { // Only go to unvisited nodes
				visited[neighbor] = 1
				traverse(neighbor) // Recursively dive into this branch
			}
		}
	}

	// Start DFS from the given starting node
	traverse(start)

	return ans
}
