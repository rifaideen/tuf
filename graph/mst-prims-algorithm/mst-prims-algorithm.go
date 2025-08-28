package main

type Edge struct {
	node   int
	weight int
}

func spanningTree(v int, edges [][]int) int {
	total := 0 // This will accumulate the total weight of the MST

	// Track which nodes have been visited (included in the MST)
	visited := make([]int, v)

	// Build an adjacency list to represent the graph
	adjList := make([][]Edge, v)
	for _, edge := range edges {
		u, v, weight := edge[0], edge[1], edge[2]

		// Add edge from u -> v and v -> u (undirected graph)
		adjList[u] = append(adjList[u], Edge{
			node:   v,
			weight: weight,
		})
		adjList[v] = append(adjList[v], Edge{
			node:   u,
			weight: weight,
		})
	}

	// Use a min-heap to always pick the smallest edge weight available
	heap := &MinHeap{
		items: []*Pair{},
	}

	// Start from node 0 with weight 0 (acts as the first edge to kick off the algorithm)
	heap.Push(&Pair{
		node:   0,
		weight: 0,
	})

	// Continue until all reachable nodes are processed
	for heap.size > 0 {
		// Extract the node with the minimum edge weight
		node := heap.Pop()

		// If this node is already in the MST, skip it
		if visited[node.node] == 1 {
			continue
		}

		// Mark node as visited (now part of MST) and add its weight to total
		visited[node.node] = 1
		total += node.weight

		// Explore all neighbors of the current node
		for _, neighbor := range adjList[node.node] {
			// Skip if neighbor is already in the MST
			if visited[neighbor.node] == 1 {
				continue
			}

			// Otherwise, add the edge weight to the heap for future consideration
			// This edge could connect to the MST with minimal cost later
			heap.Push(&Pair{
				node:   neighbor.node,
				weight: neighbor.weight,
			})
		}
	}

	// Return the total weight of the minimum spanning tree
	return total
}

type Pair struct {
	node   int
	weight int
}

type MinHeap struct {
	items []*Pair
	size  int
}

func (h *MinHeap) Push(value *Pair) {
	h.items = append(h.items, value)
	h.Up(h.size)
	h.size++
}

func (h *MinHeap) Pop() *Pair {
	if h.size == 0 {
		return nil
	}

	top := h.items[0]
	h.size--
	h.items[0] = h.items[h.size]
	h.items = h.items[:h.size]

	h.Down(0)

	return top
}

func (h *MinHeap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.items[index].weight < h.items[parent].weight {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MinHeap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1
	smallest := index

	if left < h.size && h.items[left].weight < h.items[smallest].weight {
		smallest = left
	}

	if right < h.size && h.items[right].weight < h.items[smallest].weight {
		smallest = right
	}

	if index != smallest {
		h.Swap(index, smallest)
		h.Down(smallest)
	}
}

func (h *MinHeap) Swap(a, b int) {
	h.items[a], h.items[b] = h.items[b], h.items[a]
}
