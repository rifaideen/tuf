package main

func minCost(ropes []int) int {
	h := &Minheap{
		nums: []int{},
		size: 0,
	}

	for _, rope := range ropes {
		h.Push(rope)
	}

	// prefix sum
	prefix_sum := 0

	for h.size > 1 {
		v1 := h.Pop()
		v2 := h.Pop()

		sum := v1 + v2
		prefix_sum += sum

		h.Push(sum)
	}

	return prefix_sum
}

type Minheap struct {
	nums []int
	size int
}

func (h *Minheap) Push(value int) {
	h.nums = append(h.nums, value)
	h.Up(h.size)

	h.size++
}

func (h *Minheap) Pop() int {
	if h.size == 0 {
		panic("no values found")
	}

	top := h.nums[0]

	h.size--
	h.nums[0] = h.nums[h.size]
	h.nums = h.nums[:h.size]
	h.Down(0)

	return top
}

func (h *Minheap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.nums[index] < h.nums[parent] {
			h.Swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *Minheap) Down(index int) {
	left := 2*index + 1
	right := left + 1
	smallest := index

	if left < h.size && h.nums[left] < h.nums[smallest] {
		smallest = left
	}

	if right < h.size && h.nums[right] < h.nums[smallest] {
		smallest = right
	}

	if index != smallest {
		h.Swap(index, smallest)
		h.Down(smallest)
	}
}

func (h *Minheap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}
