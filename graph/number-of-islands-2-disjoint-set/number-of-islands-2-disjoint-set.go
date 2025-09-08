package main

// numIslands simulates adding land cells one by one on an initially empty grid,
// and returns the number of connected islands after each operation.
//
// Uses Union-Find (Disjoint Set Union, DSU) to dynamically track connected components.
// An island is a group of 1's connected 4-directionally (up/down/left/right).
//
// Intuition:
// - Start with 0 islands.
// - For each new land cell:
//  1. If already added, just record current island count.
//  2. Otherwise, add it and increment island count (tentatively).
//  3. Check its 4 neighbors: if any is land, try to merge.
//  4. Each successful merge (different component) reduces island count by 1.
//
// - Result: incremental island count after each operation.
func numIslands(rows, cols int, operators [][]int) []int {
	// visited[row][col] = 1 if land has been placed at (row, col)
	visited := make([][]int, rows)

	// Initialize each row to have 'cols' columns
	for row := range rows {
		visited[row] = make([]int, cols)
	}

	// DSU to manage connected components; each cell is a node: index = row*cols + col
	ds := NewDS(rows * cols)
	counter := 0   // current number of islands
	ans := []int{} // result: island count after each operation

	// Directions for 4 adjacent cells: up, down, left, right
	delRows := []int{-1, 1, 0, 0}
	delCols := []int{0, 0, -1, 1}

	// Process each land addition
	for _, node := range operators {
		row := node[0]
		col := node[1]

		// If this cell is already land, island count doesn't change
		if visited[row][col] == 1 {
			ans = append(ans, counter)
			continue
		}

		// Mark as land and increment island count (new isolated island)
		visited[row][col] = 1
		counter++

		// Explore all 4 neighboring cells
		for i := range 4 {
			r := row + delRows[i]
			c := col + delCols[i]

			// Check if neighbor is within grid bounds
			if r >= 0 && r < rows && c >= 0 && c < cols {
				// Only consider if neighbor is already land
				if visited[r][c] == 1 {
					// Convert 2D coordinates to 1D index for DSU
					nodeIndex := row*cols + col // current cell
					adjNodeIndex := r*cols + c  // adjacent land cell

					// Find ultimate parents (components) of both cells
					parent1 := ds.FindUltimateParent(nodeIndex)
					parent2 := ds.FindUltimateParent(adjNodeIndex)

					// If they're in different components, merge them
					if parent1 != parent2 {
						ds.UnionByRank(nodeIndex, adjNodeIndex)
						counter-- // two islands become one
					}
				}
			}
		}

		// Record island count after processing this operation
		ans = append(ans, counter)
	}

	return ans
}

// DisjointSet (Union-Find) data structure to manage connected components.
// Supports:
// - FindUltimateParent: find root of a node (with path compression)
// - UnionByRank: merge two sets by rank to keep tree flat
// - UnionBySize: alternative union strategy (not used here)
type DisjointSet struct {
	parents []int // parents[i] = parent of node i
	sizes   []int // size of component rooted at i (used in UnionBySize)
	ranks   []int // rank (approx depth) of node i (used in UnionByRank)
}

// NewDS initializes a DisjointSet with n nodes.
// Each node starts as its own parent (self-loop), size 1, rank 0.
// This is essential for DSU to work correctly.
func NewDS(n int) *DisjointSet {
	parents := make([]int, n)
	sizes := make([]int, n)
	ranks := make([]int, n)

	// Initialize each node to point to itself
	for i := range n {
		parents[i] = i
		sizes[i] = 1
		ranks[i] = 0
	}

	return &DisjointSet{
		parents: parents,
		sizes:   sizes,
		ranks:   ranks,
	}
}

// FindUltimateParent finds the root of the set containing 'node'.
// Uses path compression: after finding the root, update parent pointers
// to point directly to the root. This flattens the tree for future queries.
// Time: nearly O(1) amortized.
func (d *DisjointSet) FindUltimateParent(node int) int {
	if node == d.parents[node] {
		return node // root found
	}

	// Recursively find root and compress path
	d.parents[node] = d.FindUltimateParent(d.parents[node])
	return d.parents[node]
}

// UnionByRank merges two sets by comparing their ranks.
// The tree with smaller rank is attached under the one with larger rank.
// If ranks are equal, one is chosen as parent and its rank is incremented.
// This prevents tree from becoming too deep.
func (d *DisjointSet) UnionByRank(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	// If already in same component, no need to union
	if pu == pv {
		return
	}

	// Attach smaller rank tree under larger rank tree
	if d.ranks[pu] < d.ranks[pv] {
		d.parents[pu] = pv
	} else if d.ranks[pv] < d.ranks[pu] {
		d.parents[pv] = pu
	} else {
		// Ranks are equal: attach one to the other and increase parent's rank
		d.parents[pv] = pu
		d.ranks[pu]++
	}
}

// UnionBySize merges two sets by comparing their sizes.
// The smaller component is attached to the larger one.
// Size of the new root is updated accordingly.
// Not used in this solution, but included for completeness.
func (d *DisjointSet) UnionBySize(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)

	if pu == pv {
		return
	}

	// Attach smaller component to larger one
	if d.sizes[pu] < d.sizes[pv] {
		d.parents[pu] = pv
		d.sizes[pv] += d.sizes[pu]
	} else {
		d.parents[pv] = pu
		d.sizes[pu] += d.sizes[pv]
	}
}
