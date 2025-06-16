package main

// Replace the elements by it's rank. When there are duplicate elements, it will be on the same rank
//
// The problem can be solved by utilizing the modified min-heap approach.
// Inserting the number and it's current index to the heap
//
// Once done, we start the pop them from the heap and keep the rank in map data structure and update the nums array
func replaceByRank(nums []int) {
	h := MinHeap{
		size: 0,
		nums: make([][]int, 0),
	}

	// push the number and it's corresponding index
	for index, num := range nums {
		h.push([]int{num, index})
	}

	// keep track of current rank and rank map
	rank := 0
	ranking := map[int]int{}

	for !h.empty() {
		val := h.pop()

		// replace the rank only for the first time and ignore duplicates
		if _, ok := ranking[val[0]]; !ok {
			rank++
			ranking[val[0]] = rank
		}

		// we know the rank of popped value and it's original index, replace them
		nums[val[1]] = rank
	}
}

// MinHeap represent the min-heap data structure with list of ListNode
type MinHeap struct {
	size int
	nums [][]int // value, current index
}

func (h *MinHeap) push(num []int) {
	h.nums = append(h.nums, num)
	h.up(h.size)
	h.size++
}

func (h *MinHeap) pop() []int {
	if h.size == 0 {
		panic("heap empty")
	}

	num := h.nums[0]

	h.size--
	h.nums[0] = h.nums[h.size]
	h.nums = h.nums[:h.size]

	h.down(0)

	return num
}

func (h *MinHeap) up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		// we take advantage of min-heap data structure of top being the smallest
		// compare current and parent value
		if h.nums[index][0] < h.nums[parent][0] {
			h.swap(index, parent)
			h.up(parent)
		}
	}
}

func (h *MinHeap) down(index int) {
	left := 2*index + 1
	right := left + 1
	smallest := index

	// compare current and left node value
	if left < h.size && h.nums[left][0] < h.nums[smallest][0] {
		smallest = left
	}

	// compare current and right node value
	if right < h.size && h.nums[right][0] < h.nums[smallest][0] {
		smallest = right
	}

	if index != smallest {
		h.swap(index, smallest)
		h.down(smallest)
	}
}

func (h *MinHeap) swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *MinHeap) empty() bool {
	return h.size == 0
}
