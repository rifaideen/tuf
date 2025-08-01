package main

// findCircleNum returns the number of provinces in a given adjacency matrix.
// A province is a group of cities directly or indirectly connected, forming a connected component.
// We treat the matrix as an undirected graph where each city is a node,
// and an edge exists between two cities if isConnected[i][j] == 1.
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected) // Number of cities (nodes)

	// Convert the adjacency matrix into an adjacency list for easier traversal
	list := build(isConnected)

	// visited[i] tracks whether city i has been visited during DFS
	visited := make([]int, n)

	provinces := 0 // Count of connected components (i.e., provinces)

	// Define DFS as a closure so it can access 'visited' and 'list' without passing them repeatedly
	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = 1 // Mark current city as visited

		// Explore all directly connected cities (neighbors)
		for _, neighbor := range list[node] {
			if visited[neighbor] == 0 { // If neighbor hasn't been visited
				dfs(neighbor) // Recursively visit it
			}
		}
	}

	// Iterate over all cities
	for i := range n {
		if visited[i] == 0 { // If city i is not yet visited, it's a new province
			provinces++
			dfs(i) // Use DFS to mark all cities in this province as visited
		}
	}

	return provinces
}

// build converts the adjacency matrix into an adjacency list representation.
// This makes graph traversal (like DFS) more efficient by avoiding repeated matrix scans.
func build(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	// Initialize adjacency list with empty slices for each node
	adjacencyList := make([][]int, rows)

	// For each pair of cities, check if they are connected
	for row := range rows {
		for col := range cols {
			// Skip self-loops (a city is always connected to itself in matrix but not considered an edge)
			// Only add an edge if there is a connection (matrix[row][col] == 1)
			if row != col && matrix[row][col] == 1 {
				adjacencyList[row] = append(adjacencyList[row], col)
			}
		}
	}

	return adjacencyList
}
