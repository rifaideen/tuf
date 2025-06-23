package main

func topKFrequent(nums []int, k int) []int {
	// create map to store the nums frequency
	frequency := map[int]int{}

	for _, num := range nums {
		frequency[num]++
	}

	// create max heap to push the value and it's frequency
	h := MaxHeap{
		size:   0,
		values: [][]int{},
	}

	for key, val := range frequency {
		h.Push([]int{val, key})
	}

	ans := []int{}

	// pop the top k values and push the value to the answer
	for range k {
		ans = append(ans, h.Pop()[1])
	}

	return ans
}

type MaxHeap struct {
	size   int
	values [][]int
}

func (h *MaxHeap) Push(value []int) {
	h.values = append(h.values, value)
	h.Up(h.size)

	h.size++
}

func (h *MaxHeap) Pop() []int {
	top := h.values[0]

	h.size--
	h.values[0] = h.values[h.size]
	h.values = h.values[:h.size]

	h.Down(0)

	return top
}

func (h *MaxHeap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.values[index][0] > h.values[parent][0] {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *MaxHeap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1
	largest := index

	if left < h.size && h.values[left][0] > h.values[largest][0] {
		largest = left
	}

	if right < h.size && h.values[right][0] > h.values[largest][0] {
		largest = right
	}

	if index != largest {
		h.Swap(index, largest)
		h.Down(largest)
	}
}

func (h *MaxHeap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
