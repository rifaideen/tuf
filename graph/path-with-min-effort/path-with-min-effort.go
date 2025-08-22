package main

import "math"

type Pair struct {
	distance int
	row      int
	col      int
}

func minimumEffortPath(heights [][]int) int {
	rows := len(heights)
	cols := len(heights[0])

	// We'll keep track of the minimum "effort" (maximum absolute height difference)
	// required to reach each cell from the starting point (0, 0).
	distances := make([][]int, rows)

	// Initialize all distances to a very large value (infinity equivalent),
	// because we haven't explored any paths yet.
	for row := range rows {
		distances[row] = make([]int, cols)
		for col := range cols {
			distances[row][col] = 1e9
		}
	}

	// The effort to reach the starting cell (0, 0) is zero.
	distances[0][0] = 0

	// Min-heap (priority queue) to always process the cell with the smallest effort first.
	// This is key to Dijkstra's algorithm: greedy selection of the next best option.
	heap := &MinHeap{
		items: []*Pair{},
	}

	// Start from the top-left corner (0, 0) with 0 effort.
	heap.Push(&Pair{
		distance: 0,
		row:      0,
		col:      0,
	})

	// Continue until all reachable cells are processed or we reach the destination.
	for heap.size > 0 {
		// Get the cell with the smallest current effort.
		// This ensures we're always expanding the least-effort path first.
		node := heap.Pop()

		// If we've reached the bottom-right corner, return its effort.
		// Since we use a min-heap, the first time we pop this node is the minimum effort!
		if node.row == rows-1 && node.col == cols-1 {
			return node.distance
		}

		// Directions: up, down, left, right
		delRows := []int{-1, 1, 0, 0}
		delCols := []int{0, 0, -1, 1}

		// Explore all 4 neighboring cells
		for i := range 4 {
			row := node.row + delRows[i]
			col := node.col + delCols[i]

			// Check if the neighbor is within grid bounds
			if row >= 0 && row < rows && col >= 0 && col < cols {
				// Calculate the height difference between current and neighbor cell
				diff := heights[node.row][node.col] - heights[row][col]
				// The effort to reach the neighbor is the maximum of:
				// - the effort it took to reach the current cell (node.distance)
				// - the absolute height difference to this neighbor
				distance := max(node.distance, int(math.Abs(float64(diff))))

				// If we found a path to this neighbor with *less* effort,
				// update its effort and add it to the heap for future exploration.
				if distance < distances[row][col] {
					distances[row][col] = distance

					heap.Push(&Pair{
						distance: distance,
						row:      row,
						col:      col,
					})
				}
			}
		}
	}

	// If no path was found (shouldn't happen in a connected grid), return 0.
	return 0
}

type MinHeap struct {
	size  int
	items []*Pair
}

func (h *MinHeap) Push(pair *Pair) {
	h.items = append(h.items, pair)
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

		if h.items[index].distance < h.items[parent].distance {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MinHeap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1

	smallest := index

	if left < h.size && h.items[left].distance < h.items[smallest].distance {
		smallest = left
	}

	if right < h.size && h.items[right].distance < h.items[smallest].distance {
		smallest = right
	}

	if smallest != index {
		h.Swap(smallest, index)
		h.Down(smallest)
	}
}

func (h *MinHeap) Swap(a, b int) {
	h.items[a], h.items[b] = h.items[b], h.items[a]
}
