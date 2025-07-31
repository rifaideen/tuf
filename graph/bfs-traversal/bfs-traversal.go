package main

// createGraph creates an adjacency list representation of the graph from the given edge list.
// Each element in the graph represents a node. The value associated with each node is a list of
// its neighbors. The graph is 1-indexed.
func createGraph(n int, list [][]int) [][]int {
	graph := make([][]int, n+1)

	// Iterate through the edge list and add edges to the graph.
	for _, data := range list {
		i, j := data[0], data[1]

		// Add j as a neighbor of i and i as a neighbor of j (undirected graph).
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}

	return graph
}

// bfsTraversal performs a Breadth-First Search (BFS) traversal of the graph starting from the given start node.
// It returns a list of nodes visited in BFS order.
func bfsTraversal(n, start int, list [][]int) []int {
	visited := make([]int, n+1)
	graph := createGraph(n, list)

	// Initialize a queue for BFS.
	queue := []int{
		start,
	}
	// Mark the starting node as visited.
	visited[start] = 1

	// Initialize a list to store the BFS traversal order.
	ans := []int{}

	// Continue the traversal as long as the queue is not empty.
	for len(queue) > 0 {
		// Dequeue a node from the front of the queue.
		node := queue[0]
		queue = queue[1:]
		// Add the dequeued node to the traversal order.
		ans = append(ans, node)

		// Mark the current node as visited
		visited[node] = 1

		// Iterate through the neighbors of the current node.
		for _, neighbor := range graph[node] {
			// If a neighbor has not been visited, mark it as visited and enqueue it.
			if visited[neighbor] == 0 {
				visited[neighbor] = 1
				queue = append(queue, neighbor)
			}
		}
	}

	return ans
}
