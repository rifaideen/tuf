package main

// Edge represents a directed edge from one node to another with a given weight (time delay)
type Edge struct {
	node   int // destination node
	weight int // time it takes to traverse this edge
}

// Pair is used in the min-heap to prioritize nodes by shortest known distance
type Pair struct {
	distance int // current shortest distance to this node
	node     int // node identifier
}

// networkDelayTime calculates the minimum time for a signal to reach all nodes from node k
// Returns -1 if it's impossible for all nodes to receive the signal
func networkDelayTime(times [][]int, n int, k int) int {
	// Step 1: Build adjacency list representation of the graph
	// adjList[i] will contain all edges going out from node i
	adjList := make([][]Edge, n+1) // index 0 unused; nodes are 1-indexed

	// distances[i] will store the shortest known distance from node k to node i
	distances := make([]int, n+1)

	// Initialize all distances to infinity (a very large number), except source
	for i := range distances {
		distances[i] = 1e9
	}
	distances[k] = 0 // distance to source node is zero

	// Populate adjacency list from input edges
	for _, time := range times {
		u, v, w := time[0], time[1], time[2] // from u -> v with weight w
		adjList[u] = append(adjList[u], Edge{
			node:   v,
			weight: w,
		})
	}

	// Step 2: Use Dijkstra's algorithm with a min-heap to find shortest paths
	// Min-heap will always give us the node with the smallest known distance
	heap := &MinHeap{items: []*Pair{}}

	// Start with the source node k at distance 0
	heap.Push(&Pair{
		node:     k,
		distance: 0,
	})

	// Process nodes in order of increasing distance (greedy choice)
	for heap.size > 0 {
		// Get the node with the smallest current distance
		curr := heap.Pop()

		// Early optimization: if we've already found a better path to this node, skip
		// (This handles duplicate entries in heap due to pushes before updates)
		if curr.distance > distances[curr.node] {
			continue
		}

		// Explore all neighbors of the current node
		for _, neighbor := range adjList[curr.node] {
			// Calculate new distance to neighbor via current node
			newDist := curr.distance + neighbor.weight

			// If we found a shorter path to the neighbor, update it
			if newDist < distances[neighbor.node] {
				distances[neighbor.node] = newDist

				// Push the updated distance into the heap
				// Note: We don't remove old entries; we'll just skip them later using the distance check
				heap.Push(&Pair{
					node:     neighbor.node,
					distance: newDist,
				})
			}
		}
	}

	// Step 3: After Dijkstra, check if all nodes are reachable and find max delay
	ans := 0
	for i := 1; i <= n; i++ {
		// If any node is still at infinite distance, it's unreachable
		if distances[i] == 1e9 {
			return -1
		}
		// The answer is the maximum shortest distance from k to any node
		// This represents the time when the last node receives the signal
		ans = max(ans, distances[i])
	}

	return ans
}

// MinHeap is a priority queue that keeps the smallest distance node on top
type MinHeap struct {
	size  int     // number of elements in the heap
	items []*Pair // slice representing the binary heap
}

// Push adds a new pair to the heap and restores heap property upward
func (h *MinHeap) Push(pair *Pair) {
	h.items = append(h.items, pair)
	h.Up(h.size) // bubble up the newly added item
	h.size++
}

// Pop removes and returns the minimum element (top of heap)
func (h *MinHeap) Pop() *Pair {
	if h.size == 0 {
		return nil
	}

	top := h.items[0]            // minimum element
	h.size--                     // decrease size
	h.items[0] = h.items[h.size] // move last element to root
	h.items = h.items[:h.size]   // trim slice

	h.Down(0) // restore heap property by bubbling down new root
	return top
}

// Up moves the element at index up the tree while it's smaller than its parent
func (h *MinHeap) Up(index int) {
	if index == 0 {
		return // root can't go up
	}
	parent := (index - 1) / 2
	// If current node is smaller than parent, swap and continue up
	if h.items[index].distance < h.items[parent].distance {
		h.Swap(index, parent)
		h.Up(parent)
	}
}

// Down moves the element at index down the tree, maintaining heap property
func (h *MinHeap) Down(index int) {
	left := 2*index + 1
	right := left + 1
	smallest := index

	// Find the smallest among node, left child, and right child
	if left < h.size && h.items[left].distance < h.items[smallest].distance {
		smallest = left
	}
	if right < h.size && h.items[right].distance < h.items[smallest].distance {
		smallest = right
	}

	// If smallest is not the current node, swap and continue down
	if smallest != index {
		h.Swap(smallest, index)
		h.Down(smallest)
	}
}

// Swap exchanges two elements in the heap
func (h *MinHeap) Swap(a, b int) {
	h.items[a], h.items[b] = h.items[b], h.items[a]
}
