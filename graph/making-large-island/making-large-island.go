package main

func largestIsland(grid [][]int) int {
	n := len(grid)
	ds := NewDS(n * n) // DSU to track connected land components

	delRows := []int{-1, 1, 0, 0}
	delCols := []int{0, 0, -1, 1}

	// Step 1: Union all adjacent land cells (1s) into components
	for row := range n {
		for col := range n {
			if grid[row][col] == 0 {
				continue // skip water
			}

			// Try to union current land cell with its 4 neighbors
			for i := range 4 {
				r := row + delRows[i]
				c := col + delCols[i]

				if r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 1 {
					u := r*n + c
					v := row*n + col
					ds.UnionBySize(u, v) // merge neighboring lands
				}
			}
		}
	}

	maxCount := 0

	// Step 2: Try flipping each water cell (0) to see if we can form a bigger island
	for row := range n {
		for col := range n {
			if grid[row][col] == 1 {
				continue // skip land, only consider water
			}

			components := map[int]struct{}{} // Track unique root parents of neighboring lands

			// Check all 4 neighbors
			for i := range 4 {
				r := row + delRows[i]
				c := col + delCols[i]

				if r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 1 {
					u := r*n + c
					root := ds.FindUltimateParent(u)
					components[root] = struct{}{} // avoid duplicate counting
				}
			}

			totalSize := 0
			// Sum sizes of all distinct neighboring land components
			for component := range components {
				totalSize += ds.sizes[component]
			}

			// Add 1 for the flipped cell (current 0 → 1)
			maxCount = max(maxCount, totalSize+1)
		}
	}

	// Step 3: Handle case where grid is all land (no 0 to flip), so answer is largest existing island
	for i := range n * n {
		root := ds.FindUltimateParent(i)
		maxCount = max(maxCount, ds.sizes[root])
	}

	return maxCount
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
