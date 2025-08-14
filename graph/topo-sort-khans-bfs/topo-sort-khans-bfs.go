package main

// topoSort returns topological ordering of vertices using Kahn's algorithm (BFS-based)
// Intuition:
//
//	We progressively remove "source" nodes (with no incoming edges) since they must come first.
//	Each removal reduces dependencies for neighboring nodes, potentially creating new sources.
func topoSort(vertices int, adjList [][]int) []int {
	ans := []int{}
	// Track how many edges point TO each node (prerequisites needed before processing)
	inDegree := make([]int, vertices)

	// Calculate in-degree for every node
	// Why? We need to know which nodes have no prerequisites (inDegree=0)
	for node := range adjList {
		for _, neighbor := range adjList[node] {
			// Skip self-loops (though topo sort assumes DAG, this is defensive)
			if node != neighbor {
				inDegree[neighbor]++
			}
		}
	}

	// Queue will hold nodes with NO incoming edges (ready to be processed)
	queue := []int{}

	// Find all initial source nodes (no prerequisites)
	// Why? These can be placed at the start of our ordering
	for i := range vertices {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Process nodes in order of dependency resolution
	// Why BFS? We process all current sources before their dependents
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Dequeue

		// This node has all prerequisites satisfied - safe to add to ordering
		ans = append(ans, node)

		// "Remove" this node from the graph:
		//   Reduce in-degree of all neighbors (since edge from node→neighbor is gone)
		for _, neighbor := range adjList[node] {
			inDegree[neighbor]--

			// If neighbor now has NO prerequisites, it becomes a new source
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Why this works:
	//   1. We only add nodes when their prerequisites are satisfied
	//   2. Each edge is processed exactly once (O(V+E) time)
	//   3. If graph has cycle, ans will have < vertices nodes (but problem assumes DAG)
	return ans
}
