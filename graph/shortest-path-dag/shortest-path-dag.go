package main

// Pair represents an edge from a node to a neighbor with a given distance.
type pair struct {
	node     int // The neighboring node
	distance int // The weight (distance) of the edge to this neighbor
}

// shortestPath computes the shortest path from node 0 to all other nodes in a
// directed, weighted graph using a topological sort (via DFS) and dynamic programming.
// This approach works only for **Directed Acyclic Graphs (DAGs)**.
//
// Input:
//
//	n: number of nodes (labeled 0 to n-1)
//	m: number of edges (unused, but provided for context)
//	edges: list of edges in the form [u, v, w] — directed edge from u to v with weight w
//
// Output:
//
//	distances[i] = shortest distance from node 0 to node i
//	If a node is unreachable, it remains at a large value (1e9), unless updated.
func shortestPath(n, m int, edges [][]int) []int {
	// visited[i] == 1 means node i has been processed in DFS (for topological sort)
	visited := make([]int, n)

	// adjList[i] = list of neighbors of node i, along with edge weights
	adjList := make([][]pair, n)

	// Build the adjacency list from the edge list
	for i := range edges {
		u, v, distance := edges[i][0], edges[i][1], edges[i][2]

		adjList[u] = append(adjList[u], pair{
			node:     v,
			distance: distance,
		})
	}

	// Stack to store nodes in topologically sorted order (via DFS finish times)
	stack := []int{}

	// distances[i] will store the shortest distance from node 0 to node i
	distances := make([]int, n)

	// Initialize all distances to a very large number (infinity)
	// Perform DFS on every unvisited node to generate topological order
	// This ensures all nodes are processed, even if the graph is disconnected
	for node := range n {
		distances[node] = 1e9

		if visited[node] == 0 {
			dfs(node, visited, &stack, adjList)
		}
	}

	// Set the source node (node 0) distance to 0
	distances[0] = 0

	// Process nodes in topological order
	// By the time we process a node, all its dependencies (inbound paths) have been considered
	for len(stack) > 0 {
		// Pop node from the end of the stack (topological order)
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Relax all outgoing edges from this node
		for _, neighbor := range adjList[node] {
			// If we can reach the neighbor faster via this node, update its distance
			newDist := distances[node] + neighbor.distance

			if newDist < distances[neighbor.node] {
				distances[neighbor.node] = newDist
			}
		}
	}

	// Convert any remaining infinity (1e9) to -1 to indicate unreachable nodes
	for node, value := range distances {
		if value == 1e9 {
			distances[node] = -1
		}
	}

	// Return the shortest distances from node 0 to all nodes
	return distances
}

// dfs performs a depth-first search to help generate a topological ordering.
// Nodes are added to the stack **after** all their descendants are processed.
// This ensures that in the stack, each node appears after all nodes it depends on.
//
// Parameters:
//
//	node: current node being visited
//	visited: tracks which nodes have been processed
//	stack: pointer to the stack to push nodes in reverse DFS finish order
//	adjList: adjacency list representing the graph
func dfs(node int, visited []int, stack *[]int, adjList [][]pair) {
	visited[node] = 1 // Mark node as visited

	// Visit all unvisited neighbors
	for _, neighbor := range adjList[node] {
		if visited[neighbor.node] == 0 {
			dfs(neighbor.node, visited, stack, adjList)
		}
	}

	// After visiting all neighbors, push the node onto the stack
	// This gives us a reverse post-order, which is a valid topological sort for DAGs
	*stack = append(*stack, node)
}
