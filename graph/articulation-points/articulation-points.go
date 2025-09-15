package main

// Global counter to assign unique discovery times during DFS traversal
var counter = 0

// Function to find the number of articulation points in an undirected graph
// An articulation point (or cut vertex) is a node whose removal increases the number of connected components
func articulationPoints(n int, edges [][]int) int {
	// visited[i] tracks whether node i has been visited in DFS
	visited := make([]int, n)
	// marked[i] will be set to 1 if node i is an articulation point
	marked := make([]int, n)
	// insertionTimes[i] records the time when node i was first discovered
	insertionTimes := make([]int, n)
	// lows[i] stores the lowest insertion time reachable from the subtree rooted at i
	lows := make([]int, n)
	// Adjacency list representation of the graph
	adjList := make([][]int, n)

	// Initialize insertionTimes and lows to -1 (unvisited state)
	for i := range n {
		lows[i] = -1
		insertionTimes[i] = -1
	}

	// Build the undirected graph: add both u->v and v->u edges
	for _, edge := range edges {
		u, v := edge[0], edge[1]

		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	// Perform DFS on each unvisited component (in case graph is disconnected)
	for i := range n {
		if visited[i] == 0 {
			dfs(i, -1, adjList, visited, insertionTimes, lows, &marked)
		}
	}

	// Count total articulation points
	ans := 0

	for i := range marked {
		ans += marked[i]
	}

	// Return the count; if no articulation points found, return -1 as per problem convention
	if ans > 0 {
		return ans
	}

	return -1
}

// DFS function to explore the graph and identify articulation points using Tarjan's algorithm
func dfs(node, parent int, adjList [][]int, visited, insertionTimes, lows []int, marked *[]int) {
	// Mark current node as visited
	visited[node] = 1
	// Assign the current counter value as the discovery (insertion) time
	insertionTimes[node] = counter
	// Initially, the lowest reachable node is itself
	lows[node] = counter
	// Increment global counter for next node
	counter++

	// Count children in DFS tree (only direct descendants from this node)
	child := 0

	// Explore all neighbors
	for _, neighbour := range adjList[node] {
		// Skip the parent to avoid going back in undirected graph
		if neighbour == parent {
			continue
		}

		// If neighbor hasn't been visited, it's a child in DFS tree
		if visited[neighbour] == 0 {
			dfs(neighbour, node, adjList, visited, insertionTimes, lows, marked)

			// After returning from DFS, update low value of current node
			// The low value can propagate up: we can reach back as far as our children can
			lows[node] = min(lows[node], lows[neighbour])

			// Condition 1: If the child cannot reach back to or above the current node,
			// then removing the current node would disconnect the child.
			// This makes the current node an articulation point (except for root).
			if lows[neighbour] >= insertionTimes[node] && parent != -1 {
				(*marked)[node] = 1
			}

			child++
		} else {
			// Back edge found: update low with the discovery time of the ancestor
			// This helps track the earliest reachable node via back edges
			lows[node] = min(lows[node], insertionTimes[neighbour])
		}
	}

	// Special case for root node: if it has more than one child in DFS tree,
	// removing it would disconnect the subtrees → so it's an articulation point
	if child > 1 && parent == -1 {
		(*marked)[node] = 1
	}
}
