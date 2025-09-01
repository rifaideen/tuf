package main

// Bellman-Ford algorithm computes the shortest paths from a single source vertex
// to all other vertices in a weighted graph, even when edges have negative weights.
// It works by relaxing all edges repeatedly, up to v-1 times (where v is the number of vertices),
// since the longest possible shortest path (without cycles) can have at most v-1 edges.
// After these relaxations, one additional pass checks for negative weight cycles:
// if any distance can still be improved, such a cycle exists, making shortest paths undefined.
// In that case, the function returns [-1]; otherwise, it returns the array of shortest distances.
func bellmanFord(v int, edges [][]int, src int) []int {
	// Create an array to store the shortest distance to each vertex
	// Initialize all distances to a very large number (effectively infinity)
	distances := make([]int, v)

	// Set every vertex's distance to a big number initially
	for node := range distances {
		distances[node] = 1e8
	}

	// The source node is at distance 0 from itself
	distances[src] = 0

	// Relax all edges v-1 times: this is enough to find the shortest path in a graph without negative cycles
	for i := 0; i < v-1; i++ {
		// Go through each edge in the graph
		for _, node := range edges {
			u, vn, distance := node[0], node[1], node[2]

			// If the starting point of this edge hasn't been reached yet, skip it
			if distances[u] == 1e8 {
				continue
			}

			// Calculate the distance to the destination through this edge
			newDistance := distances[u] + distance

			// If this path is shorter, update the shortest distance to the destination
			if newDistance < distances[vn] {
				distances[vn] = newDistance
			}
		}
	}

	// After v-1 iterations, do one more pass to check for negative cycles
	// If we can still relax any edge, then a negative weight cycle exists
	for _, node := range edges {
		u, vn, distance := node[0], node[1], node[2]

		// Skip edges starting from unreachable nodes
		if distances[u] == 1e8 {
			continue
		}

		// Check if we can relax this edge again
		newDistance := distances[u] + distance

		// If a shorter path is found now, there must be a negative cycle
		if newDistance < distances[vn] {
			return []int{-1}
		}
	}

	// Return the shortest distances from the source to all other vertices
	return distances
}
