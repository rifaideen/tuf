package main

func bellmanFord(v int, edges [][]int, src int) []int {
	distances := make([]int, v)

	for node := range distances {
		distances[node] = 1e8
	}

	distances[src] = 0

	for i := 0; i < v-1; i++ {
		for _, node := range edges {
			u, vn, distance := node[0], node[1], node[2]

			if distances[u] == 1e8 {
				continue
			}

			newDistance := distances[u] + distance

			if newDistance < distances[vn] {
				distances[vn] = newDistance
			}
		}
	}

	// n-th iteration to detect negative cycle
	for _, node := range edges {
		u, vn, distance := node[0], node[1], node[2]

		if distances[u] == 1e8 {
			continue
		}

		newDistance := distances[u] + distance

		if newDistance < distances[vn] {
			return []int{-1}
		}
	}

	return distances
}
