package main

// DisjointSet represents a union-find data structure with path compression and union by rank or size.
type DisjointSet struct {
	ranks   []int // tracks the depth of trees (for union by rank)
	sizes   []int // tracks the number of nodes in the component (for union by size)
	parents []int // stores the parent of each node
}

// New initializes a DisjointSet for n nodes (1-indexed).
func New(n int) *DisjointSet {
	parents := make([]int, n+1)
	sizes := make([]int, n+1)

	// Each node starts as its own parent
	for i := range parents {
		parents[i] = i
		sizes[i] = 1
	}

	return &DisjointSet{
		ranks:   make([]int, n+1), // all ranks start at 0
		sizes:   sizes,            // all are individual components at start
		parents: parents,
	}
}

// FindUltimateParent returns the root of the node with path compression.
func (d *DisjointSet) FindUltimateParent(node int) int {
	// If node is its own parent, it's the root
	if node == d.parents[node] {
		return node
	}

	// Path compression: flatten the tree by setting parent to the root
	d.parents[node] = d.FindUltimateParent(d.parents[node])

	return d.parents[node]
}

// UnionByRank merges two sets by rank (attach smaller depth tree under deeper one).
func (d *DisjointSet) UnionByRank(u, v int) {
	pv := d.FindUltimateParent(v)
	pu := d.FindUltimateParent(u)

	// If both nodes are already in the same set, do nothing
	if pu == pv {
		return
	}

	// Union by rank: attach smaller rank tree under higher rank tree
	if d.ranks[pu] == d.ranks[pv] {
		d.parents[pu] = pv // make pv the parent
		d.ranks[pv]++      // increase rank of the new root (pv)

	} else if d.ranks[pu] < d.ranks[pv] {
		d.parents[pu] = pv // attach pu under pv
	} else {
		d.parents[pv] = pu // attach pv under pu
	}
}

// UnionBySize merges two sets by size (attach smaller component under larger one).
func (d *DisjointSet) UnionBySize(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	// If both nodes are already in the same set, do nothing
	if pu == pv {
		return
	}

	// Union by size: attach smaller tree to larger one to minimize growth
	if d.sizes[pu] < d.sizes[pv] {
		d.parents[pu] = pv         // attach pu to pv
		d.sizes[pv] += d.sizes[pu] // pv's size grows
	} else {
		d.parents[pv] = pu         // attach pv to pu
		d.sizes[pu] += d.sizes[pv] // pu's size grows
	}
}
