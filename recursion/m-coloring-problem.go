package main

func graphColoring(graph [101][101]int, m, n int) bool {
	colors := make([]int, n)
	return graphColoringRecursive(0, m, n, graph, colors)
}

func graphColoringRecursive(node, m, n int, graph [101][101]int, colors []int) bool {
	// whenever the node reaches n, we are able to paint m colors successfully
	if node == n {
		return true
	}

	// try each colors (upto m) and see if we can successfully paint
	for color := 1; color <= m; color++ {
		// check if we can color the vertex with current color
		if canColor(node, n, graph, colors, color) {
			// apply the color and proceed
			colors[node] = color

			if graphColoringRecursive(node+1, m, n, graph, colors) {
				return true
			}

			// back-track the color we have added
			colors[node] = 0
		}
	}

	return false
}

func canColor(node, n int, graph [101][101]int, colors []int, color int) bool {
	// check if adjacent colors are same
	for i := range n {
		// ignore comparing the same node (i != node)
		// check if we have adjacent vertex and that too having same color
		// to check the vertex we can use graph[node][i] or graph[i][node] both are correct
		if i != node && graph[node][i] == 1 && colors[i] == color {
			return false
		}
	}

	// none of the vertices have same color, hence we can paint the color
	return true
}
