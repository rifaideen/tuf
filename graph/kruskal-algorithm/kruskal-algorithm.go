package main

import "sort"

// spanningTree computes the total weight of the Minimum Spanning Tree (MST)
// using Kruskal's algorithm.
//
// Intuition:
// - We want to connect all vertices with minimum total edge weight, without forming cycles.
// - Kruskal’s algorithm: pick the smallest edge that doesn’t form a cycle, repeat until done.
// - Uses Union-Find (Disjoint Set) to efficiently detect cycles.
func spanningTree(v int, adjList [][][]int) int {
	total := 0 // Accumulates total weight of MST

	// Step 1: Extract all edges from the adjacency list
	// Why? Kruskal works on a global list of edges, not per-node traversal.
	// adjList[i] = [[neighbor, weight], ...] → convert to {u, v, weight}
	edges := []*Edge{}

	for i := range v { // For each node u
		for _, node := range adjList[i] { // For each edge from u
			adjNode := node[0] // v
			weight := node[1]  // weight of edge u--v

			// Store edge: from i to adjNode with given weight
			edges = append(edges, &Edge{
				node:    i,
				adjNode: adjNode,
				weight:  weight,
			})
		}
	}

	// Step 2: Sort edges by weight in increasing order
	// Why? Kruskal is greedy: always pick the smallest available edge next.
	// This ensures we build the *minimum* spanning tree.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// Step 3: Initialize Union-Find (Disjoint Set)
	// Why? To efficiently check if adding an edge forms a cycle.
	// If two nodes are already in the same component (same ultimate parent),
	// connecting them would create a cycle → skip such edges.
	ds := NewDisjointSet(v)

	// Step 4: Process each edge in sorted order
	for _, edge := range edges {
		u, v, weight := edge.node, edge.adjNode, edge.weight

		// Check if u and v are already connected (same component)
		if ds.FindUltimateParent(u) != ds.FindUltimateParent(v) {
			// They're not connected → safe to add this edge to MST
			total += weight      // Add weight to MST total
			ds.UnionByRank(u, v) // Merge their components
		}
		// If they're already connected → skip (would form cycle)
	}

	return total // Final MST weight
}

// Edge represents a directed edge from 'node' to 'adjNode' with 'weight'
// Note: Even though graph is undirected, we store both u->v and v->u implicitly
// via adjacency list. Union-Find naturally avoids duplicates.
type Edge struct {
	node    int // source node
	adjNode int // destination node
	weight  int // edge weight
}

// DisjointSet (Union-Find) tracks connected components efficiently.
// Supports two operations:
// - FindUltimateParent: find root of a node (with path compression)
// - UnionByRank: merge two components optimally to keep tree flat
type DisjointSet struct {
	parents []int // parent of each node
	ranks   []int // rank used to decide union order (keeps tree balanced)
}

// NewDisjointSet initializes a DSU with each node as its own parent.
// Initially, each node is its own component.
func NewDisjointSet(v int) *DisjointSet {
	parents := make([]int, v)
	ranks := make([]int, v) // Initial ranks 0

	for i := range v {
		parents[i] = i // Each node points to itself
	}

	return &DisjointSet{
		parents: parents,
		ranks:   ranks,
	}
}

// FindUltimateParent finds the root of the component containing 'node'
// Uses path compression: during lookup, flatten the tree so future queries are faster.
// Example: A -> B -> C -> D becomes A -> D, B -> D, C -> D after call.
func (d *DisjointSet) FindUltimateParent(node int) int {
	if node == d.parents[node] {
		return node // Root found
	}

	// Recursively find root and compress path
	d.parents[node] = d.FindUltimateParent(d.parents[node])

	return d.parents[node]
}

// UnionByRank merges two components containing u and v
// Uses rank to decide parent: attach smaller rank tree under higher rank root.
// Why? To keep the tree shallow → faster future finds.
// If ranks are equal, pick one as parent and increase its rank.
func (d *DisjointSet) UnionByRank(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	if pu == pv {
		return // Already in same component → no union needed
	}

	// Attach lower rank tree under higher rank root
	if d.ranks[pu] == d.ranks[pv] {
		d.parents[pu] = pv // Arbitrarily attach pu under pv
		d.ranks[pv]++      // Rank increases since height increased
	} else if d.ranks[pu] < d.ranks[pv] {
		d.parents[pu] = pv // pu goes under pv
	} else {
		d.parents[pv] = pu // pv goes under pu
	}
}
