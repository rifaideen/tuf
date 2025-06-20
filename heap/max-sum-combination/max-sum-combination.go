package main

import (
	"cmp"
	"slices"
)

func maxCombinations(N, K int, a, b []int) []int {
	// sort the given arrays in descending order, so that we get the higher combination values
	slices.SortFunc(a, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	slices.SortFunc(b, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	// visited indices to avoid duplicate calculation
	visited := map[[2]int]struct{}{} // [i, j] combination

	// max heap instance
	h := MaxHeap{
		nums: [][]int{}, // [][]int{sum, i, j} store the sum, i and j
	}

	// sum the first value from a and b, push it to heap along with it's indices and mark them as visited
	h.Push([]int{a[0] + b[0], 0, 0})
	visited[[2]int{0, 0}] = struct{}{}

	ans := []int{}

	// start the loop  upto k times
	for range K {
		// pop the highest value and push the value to the answer and set the i and j
		result := h.Pop()
		ans = append(ans, result[0])
		i, j := result[1], result[2]

		// check if we have visited to avoid duplicate
		_, ok := visited[[2]int{i + 1, j}]

		// make sure we haven't visited and i is within range
		if (i+1) < N && !ok {
			// mark the i+1, j as visited and push the sum to the heap
			visited[[2]int{i + 1, j}] = struct{}{}
			h.Push([]int{a[i+1] + b[j], i + 1, j})
		}

		_, ok = visited[[2]int{i, j + 1}]

		// make sure we haven't visited and j is within range
		if (j+1) < N && !ok {
			// mark the i, j+1 as visited and push the sum to the heap
			visited[[2]int{i, j + 1}] = struct{}{}
			h.Push([]int{a[i] + b[j+1], i, j + 1})
		}
	}

	return ans
}

type MaxHeap struct {
	nums [][]int
	size int
}

func (h *MaxHeap) Push(value []int) {
	h.nums = append(h.nums, value)
	h.Up(h.size)

	h.size++
}

func (h *MaxHeap) Pop() []int {
	if h.size == 0 {
		panic("no items found")
	}

	top := h.nums[0]

	h.size--
	h.nums[0] = h.nums[h.size]
	h.nums = h.nums[:h.size]
	h.Down(0)

	return top
}

func (h *MaxHeap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.nums[index][0] > h.nums[parent][0] {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MaxHeap) Down(index int) {
	left := 2*index + 1
	right := left + 1
	largest := index

	if left < h.size && h.nums[left][0] > h.nums[largest][0] {
		largest = left
	}

	if right < h.size && h.nums[right][0] > h.nums[largest][0] {
		largest = right
	}

	if index != largest {
		h.Swap(index, largest)
		h.Down(largest)
	}
}

func (h *MaxHeap) Swap(a, b int) {
	h.nums[a], h.nums[b] = h.nums[b], h.nums[a]
}
