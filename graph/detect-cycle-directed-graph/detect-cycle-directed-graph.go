package main

// detectCycle checks whether there is a cycle in a directed graph.
// The graph is represented as an adjacency list.
// Returns true if a cycle exists, false otherwise.
func detectCycle(numCourses int, adjacencyList [][]int) bool {
	// visited[i] = 1 if node i has been visited in any DFS traversal
	// This prevents redundant exploration of already processed nodes
	visited := make([]int, numCourses)

	// pathVisited[i] = 1 if node i is currently in the recursion stack (current DFS path)
	// This helps detect back edges: if we revisit a node on the same path → cycle exists
	pathVisited := make([]int, numCourses)

	// Loop through all nodes (0 to numCourses-1), since the graph may be disconnected
	// Note: Nodes are 0-indexed, so we should start from 0, not 1
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			// Start DFS from unvisited node
			if dfs(i, visited, pathVisited, adjacencyList) {
				return true // Cycle detected during DFS from node i
			}
		}
	}

	// No cycle found in any connected component of the graph
	return false
}

// dfs performs depth-first search starting from 'node'
// Returns true if a cycle is detected in the current DFS path
func dfs(node int, visited []int, pathVisited []int, adjacencyList [][]int) bool {
	// Mark the current node as visited globally and part of the current DFS path
	visited[node] = 1
	pathVisited[node] = 1

	// Explore all neighbors (i.e., nodes directly reachable from 'node')
	for _, neighbor := range adjacencyList[node] {
		if visited[neighbor] == 0 {
			// Neighbor hasn't been visited yet → explore it recursively
			if dfs(neighbor, visited, pathVisited, adjacencyList) {
				return true // Cycle found in deeper traversal
			}
		} else if pathVisited[neighbor] == 1 {
			// Neighbor has been visited AND is still in the current DFS recursion stack
			// This means we've come back to a node on the same path → back edge → cycle!
			return true
		}
		// If neighbor is visited but NOT in pathVisited, it's from a different branch
		// → no cycle, so we continue
	}

	// Backtrack: remove the current node from the recursion stack
	// It's still marked as visited globally, but no longer part of the active path
	pathVisited[node] = 0

	// No cycle found in this DFS branch
	return false
}
