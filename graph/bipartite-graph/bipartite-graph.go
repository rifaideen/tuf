package main

// isBipartite checks if a given graph is bipartite.
// A graph is bipartite if its nodes can be divided into two independent sets
// such that every edge connects a node from one set to a node from the other.
// This is equivalent to checking if the graph can be colored using two colors
// such that no two adjacent nodes have the same color (2-coloring).
func isBipartite(graph [][]int) bool {
	// colors[i] will store the color of node i.
	// We use -1 to indicate the node hasn't been visited/colored yet.
	// We'll use 0 and 1 as the two colors.
	colors := make([]int, len(graph))

	// Initialize all nodes as uncolored (-1)
	for i := range len(graph) {
		colors[i] = -1
	}

	// We loop through all nodes because the graph might be disconnected.
	// Each unvisited node starts a new connected component.
	for i := range len(graph) {
		// If this node hasn't been colored yet, start a DFS from it.
		// We try to 2-color the connected component starting at node i.
		if colors[i] == -1 {
			// Start DFS with color 0. If this component can't be 2-colored, return false.
			if !dfs(i, 0, colors, graph) {
				return false
			}
		}
	}

	// If all components are successfully 2-colored, the graph is bipartite.
	return true
}

// dfs performs a depth-first search to attempt 2-coloring the graph.
// node: current node being visited
// color: the color to assign to this node
// colors: slice tracking the color of each node
// graph: adjacency list representation of the graph
// Returns true if the component can be 2-colored, false otherwise.
func dfs(node, color int, colors []int, graph [][]int) bool {
	// Assign the current node the given color
	colors[node] = color

	// Check all neighbors of the current node
	for _, neighbor := range graph[node] {
		// If the neighbor hasn't been visited yet, color it with the opposite color
		switch colors[neighbor] {
		case -1:
			// Flip the color: 0 becomes 1, 1 becomes 0
			newColor := 1
			if color == 0 {
				newColor = 1
			} else {
				newColor = 0
			}

			// Recursively color the neighbor and its subtree
			// If at any point coloring fails, propagate false upward
			if !dfs(neighbor, newColor, colors, graph) {
				return false
			}
		case color:
			// If the neighbor is already colored with the same color as the current node,
			// this creates a conflict — two adjacent nodes have the same color.
			// Therefore, the graph cannot be 2-colored (not bipartite).
			return false
		}
		// If neighbor has opposite color, everything is fine — no action needed.
	}

	// If we successfully colored all reachable nodes without conflict, return true.
	return true
}
