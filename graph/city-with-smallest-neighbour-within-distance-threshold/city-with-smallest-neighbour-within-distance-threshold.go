package main

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Intuition:
	// The goal is to find the "best" city. The best city is defined as the one
	// that has the fewest number of other cities reachable within a given
	// `distanceThreshold`. If there's a tie, we pick the city with the highest index.
	//
	// To solve this, we first need to know the shortest distance between every pair
	// of cities. The Floyd-Warshall algorithm is a perfect tool for this task as it
	// calculates all-pairs shortest paths.
	//
	// The algorithm proceeds in three main steps:
	// 1. Initialize a distance matrix representing the graph.
	// 2. Run the Floyd-Warshall algorithm on the matrix.
	// 3. Iterate through the results to find the city that satisfies the problem's criteria.

	// ---------------------------------------------------------------------------
	// Step 1: Initialize the adjacency matrix with distances.
	// Intuition:
	// We create a 2D matrix to store the shortest distance between cities.
	// We initialize all distances to a large value (1e9) to represent "infinity"
	// (meaning no direct path exists). The distance from a city to itself is always 0.
	// Then, we populate the matrix with the direct edge weights from the input `edges`.

	distances := make([][]int, n)

	for i := range n {
		distances[i] = make([]int, n)

		for j := range n {
			// set diagonal cells (itself to itself) to 0
			if i == j {
				distances[i][j] = 0
			} else { // else set them to infinity
				distances[i][j] = 1e9
			}
		}
	}

	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]
		distances[u][v] = weight
		distances[v][u] = weight
	}

	// ---------------------------------------------------------------------------
	// Step 2: Apply the Floyd-Warshall algorithm.
	// Intuition:
	// This triple nested loop is the core of the algorithm. It systematically
	// checks every city (`via`) as a potential intermediate stop to find a shorter
	// path between every pair of cities (`i` and `j`).
	// The key update is `distances[i][j] = min(distances[i][j], distances[i][via] + distances[via][j])`.
	// This dynamic programming approach ensures that by the end, `distances[i][j]` holds the
	// shortest path between every city `i` and city `j`.

	for via := range n {
		for i := range n {
			for j := range n {
				distances[i][j] = min(distances[i][j], distances[i][via]+distances[via][j])
			}
		}
	}

	// ---------------------------------------------------------------------------
	// Step 3: Find the city with the minimum reachable cities.
	// Intuition:
	// With the all-pairs shortest paths computed, we can now solve the main problem.
	// We iterate through each city and count how many other cities are within the
	// `distanceThreshold`.
	// We maintain two variables: `cities` to store the minimum count found so far,
	// and `cityNum` to store the corresponding city index. The condition `cityCounter <= cities`
	// correctly handles both finding a new minimum and the tie-breaking rule
	// (choosing the city with the largest index in case of a tie).

	cities := n
	cityNum := -1

	for i := range n {
		cityCounter := 0

		for j := range n {
			if distances[i][j] <= distanceThreshold {
				cityCounter++
			}
		}

		if cityCounter <= cities {
			cities = cityCounter
			cityNum = i
		}
	}

	return cityNum
}
