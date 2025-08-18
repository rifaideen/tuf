package main

import "slices"

// pair represents a node and its distance
type pair struct {
	node     int
	distance int
}

// shortestPath calculates the shortest path from node 1 to node n using Dijkstra's algorithm.
func shortestPath(n, m int, edges [][]int) []int {
	// Adjacency list for storing graph connections
	adjList := make([][]pair, n+1)

	// Iterate through all provided edges
	for _, edge := range edges {
		// Extract nodes and distance from the edge
		u, v, distance := edge[0], edge[1], edge[2]
		// Append neighbor node and its distance to first node
		adjList[u] = append(adjList[u], pair{node: v, distance: distance})
		// Append neighbor node and its distance to second node (for undirected graph)
		adjList[v] = append(adjList[v], pair{node: u, distance: distance})
	}

	// Array to store distances from source node to all other nodes
	distances := make([]int, n+1)
	// Array to keep track of parent node for each node in the shortest path
	parents := make([]int, n+1)

	// Initialize all distances with a large value (infinity), and parent as themselves.
	for i := range n + 1 {
		// Large distance represents that those nodes are not yet visited.
		distances[i] = 1e9
		// Parent of each node initially set to itself, this will help in reconstructing path later on.
		parents[i] = i
	}

	// Distance from source to itself is zero.
	distances[1] = 0

	// Create a new min heap for selecting minimum weight edge efficiently during Dijkstra's algorithm execution
	h := New()
	// Insert source node into the min heap
	h.Push(pair{node: 1, distance: 0})

	// While there are still nodes with unvisited neighbors:
	for h.size > 0 {
		// Select the node with minimum distance from heap
		node := h.Pop()

		// Iterate through each neighbor of current popped node.
		for _, neighbor := range adjList[node.node] {
			// Calculate the new distance if we go though this neighbor to reach target node.
			newDistance := node.distance + neighbor.distance

			// If a shorter path is found, update the distance and parent, then insert it into priority queue
			if newDistance < distances[neighbor.node] {
				distances[neighbor.node] = newDistance
				parents[neighbor.node] = node.node

				h.Push(pair{node: neighbor.node, distance: newDistance})
			}
		}
	}

	// If destination is unreachable, because its distance didn't get updated from initial value we set for all nodes before starting Dijkstra's algorithm.
	if distances[n] == 1e9 {
		// Return -1 to denote that no path was found.
		return []int{-1}
	}

	// Initialize a slice to store nodes in shortest path
	paths := []int{}
	// Start with destination node (n)
	current := n

	// Traverse back through parent pointers from destination to source until we reach the source itself.
	for parents[current] != current {
		// Append each visited node to path slice.
		paths = append(paths, current)
		// Move one step closer towards the source using stored parent information.
		current = parents[current]
	}

	// Append final soure node (node 1) after traversing all nodes in reverse order of path.
	paths = append(paths, 1)

	// Reverse the slice to get actual shortest path from source node 1 to node n.
	slices.Reverse(paths)

	// Return reconstructed shortedPath.
	return paths
}

// MinHeap is a structure representing a min heap which will be used for efficiently finding vertex with smallest tentative distance during Dijkstra's algorithm execution.
type MinHeap struct {
	// Number of elements in heap
	size int
	// Slice containing pairs (node, distance) arranged as heap
	items []pair
}

// Initialize an empty min heap with zero size and no items in it.
func New() *MinHeap {
	return &MinHeap{
		size:  0,
		items: []pair{},
	}
}

// Push adds a new element to the min heap. It also maintains order of nodes according to their distances from source node so that smallest can always be fetched first when required by Pop operation.
func (h *MinHeap) Push(value pair) {
	// Add new pair at last index in the slice
	h.items = append(h.items, value)
	// Move newly added item up until correct position is found within heap structure according to its distance from source node (Dijkstra's min priority queue).
	h.up(h.size)
	// Increase size of heap as we just pushed a new element.
	h.size++
}

// Pop fetches and removes first element from the min heap which represents node with smallest tentative distance. It also reorganizes remaining nodes back into valid min heap afterwards
func (h *MinHeap) Pop() pair {
	// Panic if popping an empty heap.
	if h.size == 0 {
		panic("heap is empty")
	}

	// Store first element (root) in a variable and later use this value as return for popped element
	top := h.items[0]
	// Decrease size indicating one less item in heap after popping operation is complete.
	h.size--
	// Move last element (which might not have been correct position yet) to root.
	h.items[0] = h.items[h.size]
	// Remove previous last element from slice, now extra copy of that node exists at roots index but heap size has decreased and last element has also moved there too which will be taken care off by next step down operation on that index.
	h.items = h.items[:h.size]
	// Reorder remaining elements in min heap according to Dijkstra's algorithm requirement by swapping parent with smallest child until correct position is found for root node again.
	h.down(0)

	// Return popped pair containing first element from the heap which represents vertex connected to source having minimum tentative distance at this point.
	return top
}

// Swap positions between two indices within this min heaps data structure array representation without changing their content itself.
func (h *MinHeap) Swap(a, b int) {
	// Simply exchange values at provided index locations a and b
	h.items[a], h.items[b] = h.items[b], h.items[a]
}

// Move node up in the heap according to its distance value, swap it with parent until it finds correct position (minimum distance node as per Dijkstra's criteria) among its children or becomes new root of tree (max height allowed here).
func (h *MinHeap) up(index int) {
	// Only if there exists parent
	if index > 0 {
		// Compute the index of parent node
		parent := (index - 1) / 2

		// If current node is smaller than its parent, swap them to maintain child being larger than parents property.
		if h.items[index].distance < h.items[parent].distance {
			// Perform swapping of two nodes
			h.Swap(index, parent)
			// Recursively move up to fix any violations in higher levels due to current node movement.
			h.up(parent)
		}
	}
}

// Move node down in the min heap according to its distance value, swap it with smallest child until this node achieves correct position or becomes a leaf node (no more children).
func (h *MinHeap) down(index int) {
	// Compute index for left child, this will be first half as zero based indexing
	left := (2 * index) + 1
	// Similarly compute right child, it comes after the immediate next index in sequence.
	right := left + 1
	// Consider current index of given node itself till we don't know what position does minimum distance hold in its direct children list.
	smallest := index

	// Amongst three possible candidates (current node, left child, and right child), find out which one has least distances among them so that swapping can occur to maintain integrity as per Dijkstra's constraints.
	if left < h.size && h.items[smallest].distance < h.items[left].distance {
		smallest = left
	}

	if right < h.size && h.items[smallest].distance < h.items[right].distance {
		smallest = right
	}

	// If current node was the one holding minimum distance so far, then all its children must satisfy min heap property and no further swaps are necessary.
	// However if either left or/and right child had smaller distance values, do appropriate swapping now to place correct nodes at their positions.
	if smallest != index {
		h.Swap(index, smallest)
		// Also recursively apply same downward shift logic to newly inserted children to keep maintaining min heap property in remaining subtree rooted here as well.
		h.down(smallest)
	}
}
