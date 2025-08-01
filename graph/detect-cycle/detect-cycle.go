package main

// detectCycle checks whether an undirected graph contains a cycle.
// Uses BFS (via helper function) on each connected component.
//
// Parameters:
//   - v: number of vertices (1-indexed, vertices are 1 to v)
//   - e: number of edges (not used here, but provided)
//   - list: adjacency list representation of the graph (1-indexed, list[0] unused)
//
// Returns:
//   - true if a cycle exists in any connected component; false otherwise.
func detectCycle(v, e int, list [][]int) bool {
	// visited[i] == 1 means node i has been explored
	// We use a simple 0/1 flag array to track visited nodes
	visited := make([]int, v+1) // index 0 unused; nodes are 1..v

	// Loop through all vertices to handle disconnected graphs
	for i := 1; i <= v; i++ {
		// Only start BFS from unvisited nodes (new connected components)
		if visited[i] == 0 {
			// Try to detect a cycle starting from node i
			if detect(i, list, visited) {
				return true // early exit: cycle found
			}
		}
	}

	return false // no cycle found in any component
}

// detect performs BFS from a given starting node to check for cycles
// in the connected component containing 'start'.
//
// Key idea: In an undirected graph, a cycle exists if we encounter a node
// that has already been visited AND it is not the immediate parent (back edge).
//
// Parameters:
//   - start: starting node for BFS
//   - list: adjacency list
//   - visited: shared visited array to track explored nodes
//
// Returns:
//   - true if a cycle is detected during traversal; false otherwise.
func detect(start int, list [][]int, visited []int) bool {
	// Mark the starting node as visited
	visited[start] = 1

	// Queue stores pairs: [current_node, parent_node]
	// Parent is used to avoid false cycle detection (don't count back edge to parent)
	queue := [][]int{{start, -1}} // -1 indicates no parent (root node)

	// BFS loop
	for len(queue) > 0 {
		// Dequeue the front element
		front := queue[0]
		queue = queue[1:]
		node := front[0]
		parent := front[1]

		// Explore all neighbors of the current node
		for _, neighbor := range list[node] {
			if visited[neighbor] == 0 {
				// Neighbor not visited → mark it and enqueue
				visited[neighbor] = 1
				queue = append(queue, []int{neighbor, node}) // node becomes parent of neighbor
			} else if neighbor != parent {
				// Neighbor is already visited, and it's NOT the parent
				// → This means we've found a back edge to a previously visited node
				// → This is a cycle!
				return true
			}
			// If neighbor == parent, we ignore (it's the way we came)
		}
	}

	// No cycle found in this connected component
	return false
}

// detectCycle checks whether an undirected graph contains a cycle.
// Uses BFS (via helper function) on each connected component.
//
// Parameters:
//   - v: number of vertices (1-indexed, vertices are 1 to v)
//   - e: number of edges (not used here, but provided)
//   - list: adjacency list representation of the graph (1-indexed, list[0] unused)
//
// Returns:
//   - true if a cycle exists in any connected component; false otherwise.
func detectCycleDFS(v, e int, list [][]int) bool {
	// visited[i] == 1 means node i has been explored
	// We use a simple 0/1 flag array to track visited nodes
	visited := make([]int, v+1) // index 0 unused; nodes are 1..v

	// Loop through all vertices to handle disconnected graphs
	for i := 1; i <= v; i++ {
		// Only start BFS from unvisited nodes (new connected components)
		if visited[i] == 0 {
			// Try to detect a cycle starting from node i
			if dfs(i, -1, list, visited) {
				return true // early exit: cycle found
			}
		}
	}

	return false // no cycle found in any component
}

// dfs checks if a cycle exists in an undirected graph starting from 'start'.
// Uses recursive DFS and tracks the parent to avoid false cycle detection.
//
// Parameters:
//   - start: current node being explored
//   - parent: the node we came from (to avoid going back immediately)
//   - list: adjacency list representation of the graph
//   - visited: slice to track which nodes have been visited (0 = unvisited, 1 = visited)
//
// Returns:
//   - true if a cycle is found in the connected component; false otherwise.
func dfs(start, parent int, list [][]int, visited []int) bool {
	// Mark the current node as visited
	// This prevents reprocessing and helps detect back edges
	visited[start] = 1

	// Explore all neighboring nodes
	for _, neighbor := range list[start] {
		if visited[neighbor] == 0 {
			// Neighbor hasn't been visited yet → explore it recursively
			// We pass 'start' as the parent of 'neighbor' to remember where we came from
			if dfs(neighbor, start, list, visited) {
				// If the recursive call finds a cycle, propagate it upward
				return true
			}
		} else if neighbor != parent {
			// Neighbor is already visited, BUT it's NOT the parent
			// This means we've found a "back edge" to a previously visited node
			// → This is not just a reverse of the edge we came from
			// → So, there must be another path to this node → CYCLE DETECTED!
			return true
		}
		// If neighbor == parent, we ignore it
		// (It's just the edge we came from — not a cycle)
	}

	// No cycle found in this branch of traversal
	return false
}
