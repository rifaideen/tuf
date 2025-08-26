package main

func shortedDistance(matrix *[][]int) {
	n := len(*matrix)

	// Intuition:
	// The Floyd-Warshall algorithm finds the shortest paths between all pairs of
	// vertices in a weighted graph. It works by progressively considering all
	// possible intermediate vertices for each path. This function assumes the
	// input matrix represents an adjacency matrix of a directed graph. The
	// value matrix[i][j] is the weight of the edge from vertex i to vertex j.

	// ---------------------------------------------------------------------------
	// Step 1: Pre-processing the matrix to prepare for the algorithm.
	// Intuition:
	// The algorithm needs a way to represent unreachable paths. Using -1 as a
	// sentinel value is common, but it can interfere with calculations if edge
	// weights are negative. A large value, like "infinity" (1e9), is a safer
	// alternative to represent unreachable paths. The distance from a vertex to
	// itself is always 0.

	for i := range n {
		for j := range n {
			// If the cell value is -1, it means there's no direct path. We
			// replace it with a large number (1e9) to signify "infinity."
			if (*matrix)[i][j] == -1 {
				(*matrix)[i][j] = 1e9
			}

			// The shortest distance from a node to itself is always 0.
			if i == j {
				(*matrix)[i][j] = 0
			}
		}
	}

	// ---------------------------------------------------------------------------
	// Step 2: Applying the Floyd-Warshall algorithm.
	// Intuition:
	// This triple nested loop is the core of the algorithm.
	// `via` represents the **intermediate** vertex.
	// `i` represents the **source** vertex.
	// `j` represents the **destination** vertex.
	// For every pair of source (i) and destination (j) vertices, we check if
	// a path through an intermediate vertex (`via`) is shorter than the current
	// shortest path. This is a dynamic programming approach, where we build up
	// the solution by considering increasingly larger sets of intermediate vertices.

	// The algorithm:
	// For each pair of vertices (i, j), we check if a path from i to j through
	// a third vertex 'via' is shorter than the existing path from i to j.
	// `matrix[i][j]` = min(`matrix[i][j]`, `matrix[i][via]` + `matrix[via][j]`)
	// The order of the loops is crucial: the `via` loop must be the outermost.

	for via := range n {
		for i := range n {
			for j := range n {
				(*matrix)[i][j] = min((*matrix)[i][j], (*matrix)[i][via]+(*matrix)[via][j])
			}
		}
	}

	// ---------------------------------------------------------------------------
	// Step 3: Post-processing the matrix.
	// Intuition:
	// After the algorithm, any path that is still "infinity" (1e9) means there is
	// no path from the source to the destination. We convert these back to -1
	// to match the original input convention, which is a cleaner representation
	// of an unreachable path.

	for i := range n {
		for j := range n {
			if (*matrix)[i][j] == 1e9 {
				(*matrix)[i][j] = -1
			}
		}
	}
}
