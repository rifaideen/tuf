package main

// Function to find the minimum number of cable moves needed to connect all computers
// n = number of computers (labeled from 0 to n-1)
// connections = list of existing direct connections between computers
func makeConnected(n int, connections [][]int) int {
	// Step 1: Create a Disjoint Set (Union-Find) to track which computers are already connected
	ds := NewDS(n)

	// 'extra' counts how many connections are redundant (they connect already-connected computers)
	extra := 0

	// We'll count how many separate groups (components) of computers there are later
	components := 0

	// Step 2: Process each existing connection
	for _, con := range connections {
		u, v := con[0], con[1] // computers connected by this cable

		// If both computers are already in the same network (same ultimate parent),
		// this connection is unnecessary — we can remove it and use it elsewhere
		if ds.FindUltimateParent(u) == ds.FindUltimateParent(v) {
			extra++ // This cable is extra; we can reuse it
		} else {
			// Otherwise, merge the two computers into the same network
			ds.UnionByRank(u, v)
		}
	}

	// Step 3: Count how many separate networks (connected components) we have
	// Each component has one root (where parent[i] == i)
	for i := range n {
		if ds.parents[i] == i { // root node: represents a separate component
			components++
		}
	}

	// Step 4: To connect all components into one network, we need (components - 1) cables
	// If we have at least that many extra cables, we can do it
	if extra >= (components - 1) {
		return components - 1 // Minimum moves needed
	}

	// Otherwise, not enough cables — impossible to connect all
	return -1
}

// Disjoint Set (Union-Find) data structure
// Helps efficiently group nodes and find which group they belong to
type DisjointSet struct {
	parents []int // parent of each node
	ranks   []int // used to keep the tree flat (optimization)
}

// Initialize Disjoint Set with 'n' nodes
func NewDS(n int) *DisjointSet {
	parents := make([]int, n)
	ranks := make([]int, n)

	// At the start, each computer is its own parent (separate group)
	for i := range n {
		parents[i] = i
	}

	return &DisjointSet{
		parents: parents,
		ranks:   ranks,
	}
}

// Find the "ultimate parent" (group leader) of a node
// Uses path compression: flattens the structure so future queries are faster
func (d *DisjointSet) FindUltimateParent(node int) int {
	// If node is its own parent, it's the root
	if d.parents[node] == node {
		return node
	}

	// Recursively find the root, and update parent to make future lookups faster
	d.parents[node] = d.FindUltimateParent(d.parents[node])
	return d.parents[node]
}

// Union two nodes by rank (to keep tree shallow)
// Merges the groups containing u and v
func (d *DisjointSet) UnionByRank(u, v int) {
	p1 := d.FindUltimateParent(u)
	p2 := d.FindUltimateParent(v)

	// If already in same group, nothing to do
	if p1 == p2 {
		return
	}

	// Attach the smaller tree under the larger one (based on rank)
	if d.ranks[p1] < d.ranks[p2] {
		d.parents[p1] = p2
	} else if d.ranks[p2] < d.ranks[p1] {
		d.parents[p2] = p1
	} else {
		// If ranks are equal, pick one as parent and increase its rank
		d.parents[p2] = p1
		d.ranks[p1]++
	}
}
