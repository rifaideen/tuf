package main

func Represent(n, m int, list [][]int) [][]int {
	// create graph of size n + 1
	graph := make([][]int, n+1)

	for _, info := range list {
		u, v := info[0], info[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	return graph
}
