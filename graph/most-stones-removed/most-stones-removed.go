package main

// Main function: returns max stones that can be removed
func removeStones(stones [][]int) int {
	n := len(stones) // Total number of stones
	rows := 0        // Track max row index seen
	cols := 0        // Track max col index seen

	// Find the largest row and column values to size the DS
	for _, node := range stones {
		rows = max(rows, node[0]) // Update max row
		cols = max(cols, node[1]) // Update max col
	}

	ds := NewDS(rows + cols + 1) // Create DS: rows + shifted cols to avoid overlap

	stoneNodes := map[int]int{} // Tracks which nodes (row/col) actually have stones

	// Process each stone to build connections
	for _, node := range stones {
		r := node[0]            // Row index
		c := node[1] + rows + 1 // Shifted col index to prevent collision

		ds.UnionBySize(r, c) // Connect this row and column

		stoneNodes[r] = 1 // Mark row node as used
		stoneNodes[c] = 1 // Mark shifted col node as used
	}

	counter := 0 // Count of connected components

	// Count unique roots (each = one connected group)
	for key := range stoneNodes {
		if key == ds.FindUltimateParent(key) { // If it's a root
			counter++ // Then it's one group
		}
	}

	return n - counter // Can remove all except one per group
}

// Disjoint Set data structure
type DisjointSet struct {
	parents []int // Parent of each node
	ranks   []int // Rank for union by rank (unused here)
	sizes   []int // Size for union by size
}

// Initialize DS with n+1 nodes
func NewDS(n int) *DisjointSet {
	parents := make([]int, n+1)
	ranks := make([]int, n+1)
	sizes := make([]int, n+1)

	// Each node starts as its own parent
	for i := range parents {
		parents[i] = i
		sizes[i] = 1 // Each component starts with size 1
	}

	return &DisjointSet{
		parents: parents,
		ranks:   ranks,
		sizes:   sizes,
	}
}

// Find root with path compression
func (d *DisjointSet) FindUltimateParent(node int) int {
	if node == d.parents[node] { // If self is parent
		return node // Then it's the root
	}

	// Recursively find root and compress path
	d.parents[node] = d.FindUltimateParent(d.parents[node])

	return d.parents[node]
}

// Union two nodes by rank (not used in main logic)
func (d *DisjointSet) UnionByRank(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	if pu == pv {
		return // Already in same component
	}

	if d.ranks[pu] == d.ranks[pv] {
		d.parents[pu] = pv
		d.ranks[pv]++
	} else if d.ranks[pu] < d.ranks[pv] {
		d.parents[pu] = pv
	} else {
		d.parents[pv] = pu
	}
}

// Union two nodes by size (used in main function)
func (d *DisjointSet) UnionBySize(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	if pu == pv {
		return // Already connected
	}

	// Attach smaller tree under larger one
	if d.sizes[pu] < d.sizes[pv] {
		d.parents[pu] = pv
		d.sizes[pv] += d.sizes[pu] // Update size
	} else {
		d.parents[pv] = pu
		d.sizes[pu] += d.sizes[pv] // Update size
	}
}
