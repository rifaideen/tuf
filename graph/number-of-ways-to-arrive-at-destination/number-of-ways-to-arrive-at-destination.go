package main

import "math"

// Edge represents a connection from one node to another with a travel time
type Edge struct {
	node int // destination node
	time int // travel time along this edge
}

// countPaths returns the number of shortest paths from node 0 to node n-1 modulo 1e9+7
func countPaths(n int, roads [][]int) int {
	// Step 1: Build an undirected adjacency list from the roads
	adjList := make([][]Edge, n)
	// times[i] = shortest time to reach node i from node 0
	times := make([]int, n)
	// ways[i] = number of shortest paths to reach node i
	ways := make([]int, n)

	// Initialize all distances to infinity, except source (node 0)
	for i := range n {
		times[i] = math.MaxInt
		adjList[i] = []Edge{} // initialize each adjacency list
	}

	times[0] = 0 // it takes 0 time to reach the start
	ways[0] = 1  // there's exactly one way to be at the start: do nothing

	// Build bidirectional graph: roads are undirected
	for _, road := range roads {
		from, to, time := road[0], road[1], road[2]

		adjList[from] = append(adjList[from], Edge{node: to, time: time})
		adjList[to] = append(adjList[to], Edge{node: from, time: time}) // reverse edge
	}

	// Step 2: Use Dijkstra's algorithm with a min-heap to explore shortest paths
	heap := &MinHeap{items: []*Pair{}, size: 0}
	heap.Push(&Pair{node: 0, time: 0}) // start from node 0

	const mod = 1000000007 // required modulo to avoid overflow

	// Step 3: Process nodes in order of increasing shortest time
	for heap.size > 0 {
		curr := heap.Pop()

		// Optimization: skip outdated entries in the heap
		// If we've already found a better (shorter) time to curr.node, ignore this one
		if curr.time > times[curr.node] {
			continue
		}

		// Explore all neighbors of the current node
		for _, neighbor := range adjList[curr.node] {
			// Calculate new time to reach 'neighbor.node' via 'curr.node'
			newTime := curr.time + neighbor.time

			// Case 1: Found a shorter path to neighbor
			if newTime < times[neighbor.node] {
				times[neighbor.node] = newTime        // update shortest time
				ways[neighbor.node] = ways[curr.node] // inherit the number of ways

				// Push the improved path into the heap
				heap.Push(&Pair{
					node: neighbor.node,
					time: newTime,
				})

				// Case 2: Found another path with the same shortest time
			} else if newTime == times[neighbor.node] {
				// Add the number of ways from current node to total ways
				ways[neighbor.node] = (ways[neighbor.node] + ways[curr.node]) % mod
			}
			// Case 3: newTime > times[neighbor.node] → ignore (longer path)
		}
	}

	// Step 4: Return the number of shortest paths to the last node (n-1)
	return ways[n-1] % mod
}

// Pair is used in the min-heap to prioritize nodes by shortest known time
type Pair struct {
	node int
	time int
}

type MinHeap struct {
	items []*Pair
	size  int
}

func (h *MinHeap) Push(node *Pair) {
	h.items = append(h.items, node)
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

		if h.items[index].time < h.items[parent].time {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MinHeap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1

	smallest := index

	if left < h.size && h.items[left].time < h.items[smallest].time {
		smallest = left
	}

	if right < h.size && h.items[right].time < h.items[smallest].time {
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
