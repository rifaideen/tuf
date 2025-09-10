package main

import "sort"

func swimInWater(grid [][]int) int {
	n := len(grid)
	ds := NewDS(n * n) // Reuse your DSU

	// Directions for 4D movement
	delRows := []int{-1, 1, 0, 0}
	delCols := []int{0, 0, -1, 1}

	// Helper: convert (i, j) to index
	idx := func(i, j int) int {
		return i*n + j
	}

	// Step 1: Collect all cells as (elevation, i, j)
	type cell struct {
		elev, row, col int
	}
	cells := []cell{}
	for i := range n {
		for j := range n {
			cells = append(cells, cell{grid[i][j], i, j})
		}
	}

	// Sort by elevation → simulate rising water
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].elev < cells[j].elev
	})

	// Track which cells have been added (submerged)
	added := make([][]bool, n)
	for i := range added {
		added[i] = make([]bool, n)
	}

	// Step 2: Process each cell in increasing order of elevation
	for _, c := range cells {
		i, j := c.row, c.col
		added[i][j] = true // This cell is now submerged

		// Try to union with 4 neighbors if they're already added
		for d := range 4 {
			ni := i + delRows[d]
			nj := j + delCols[d]

			if ni >= 0 && ni < n && nj >= 0 && nj < n && added[ni][nj] {
				u := idx(i, j)
				v := idx(ni, nj)
				ds.UnionBySize(u, v) // Connect current cell to neighbor
			}
		}

		// Step 3: After adding this cell, check if (0,0) and (n-1,n-1) are connected
		if added[0][0] && added[n-1][n-1] { // both must be added first
			if ds.FindUltimateParent(idx(0, 0)) == ds.FindUltimateParent(idx(n-1, n-1)) {
				return c.elev // This is the minimum time t
			}
		}
	}

	return -1 // unreachable (should not happen)
}

// Disjoint Set (Union-Find) for grouping connected lands
type DisjointSet struct {
	parents []int
	ranks   []int
	sizes   []int
}

func NewDS(n int) *DisjointSet {
	parents := make([]int, n)
	ranks := make([]int, n)
	sizes := make([]int, n)

	for i := range n {
		parents[i] = i
		ranks[i] = 0
		sizes[i] = 1 // each cell starts as size 1
	}

	return &DisjointSet{
		parents: parents,
		ranks:   ranks,
		sizes:   sizes,
	}
}

// Path compression to find root parent
func (d *DisjointSet) FindUltimateParent(node int) int {
	if node == d.parents[node] {
		return node
	}
	d.parents[node] = d.FindUltimateParent(d.parents[node]) // compress path
	return d.parents[node]
}

// Union by rank (not used here, but available)
func (d *DisjointSet) UnionByRank(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)
	if pu == pv {
		return
	}
	if d.ranks[pu] < d.ranks[pv] {
		d.parents[pu] = pv
	} else if d.ranks[pv] < d.ranks[pu] {
		d.parents[pv] = pu
	} else {
		d.parents[pu] = pv
		d.ranks[pv]++
	}
}

// Union by size: attach smaller tree to larger, update size
func (d *DisjointSet) UnionBySize(u, v int) {
	pu := d.FindUltimateParent(u)
	pv := d.FindUltimateParent(v)
	if pu == pv {
		return
	}
	if d.sizes[pu] < d.sizes[pv] {
		d.parents[pu] = pv
		d.sizes[pv] += d.sizes[pu]
	} else {
		d.parents[pv] = pu
		d.sizes[pu] += d.sizes[pv]
	}
}
