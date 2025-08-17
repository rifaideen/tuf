package main

func shortestPath(edges [][]int, n, m, src int) []int {
	adjList := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	distances := make([]int, n)

	for i := range n {
		distances[i] = 1e9
	}

	distances[src] = 0
	queue := []int{src}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range adjList[node] {
			if distances[node]+1 < distances[neighbor] {
				distances[neighbor] = distances[node] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	for node, value := range distances {
		if value == 1e9 {
			distances[node] = -1
		}
	}

	return distances
}
