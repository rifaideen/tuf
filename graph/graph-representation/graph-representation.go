package main

func Represent(n, m int, list [][]int) [][]int {
	// create graph of size n + 1.  We use n+1 because graph nodes are generally 1-indexed.
	graph := make([][]int, n+1)

	// Iterate through the input list of edges. Each edge is represented as [u, v], where u and v are the nodes connected by the edge.
	for _, info := range list {
		u, v := info[0], info[1]

		// Add v to the adjacency list of u, indicating an edge from u to v.
		graph[u] = append(graph[u], v)
		// Since the graph is undirected, add u to the adjacency list of v, indicating an edge from v to u.
		graph[v] = append(graph[v], u)
	}

	return graph
}
