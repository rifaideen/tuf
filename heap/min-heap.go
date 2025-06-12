package main

type MinHeap struct {
	items []int

	size int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: []int{},
		size:  0,
	}
}

func (h *MinHeap) heapifyUp(index int) {
	parent := (index - 1) / 2

	if index > 0 && h.items[index] < h.items[parent] {
		h.items[index], h.items[parent] = h.items[parent], h.items[index]
		h.heapifyUp(parent)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	left := index*2 + 1
	right := left + 1
	smallest := index

	if left <= h.size && h.items[left] < h.items[smallest] {
		smallest = left
	}

	if right <= h.size && h.items[right] < h.items[smallest] {
		smallest = right
	}

	if index != smallest {
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]

		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) heapify() {
	for i := h.size / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MinHeap) Insert(key int) {
	h.items = append(h.items, key)
	h.heapifyUp(h.size)
	h.size++
}

func (h *MinHeap) GetMin() int {
	if h.IsEmpty() {
		return -1
	}

	return h.items[0]
}

// delete and return
func (h *MinHeap) ExtractMin() int {
	if h.IsEmpty() {
		return -1
	}

	top := h.items[0]
	h.items[0] = h.items[h.size]
	h.items = h.items[:len(h.items)-1]
	h.size--

	h.heapifyDown(0)

	return top
}

// size of heap
func (h *MinHeap) HeapSize() int {
	return h.size
}

func (h *MinHeap) IsEmpty() bool {
	return h.size == 0

}

func (h *MinHeap) ChangeKey(index, val int) {
	if val < h.items[index] {
		h.items[index] = val
		h.heapifyUp(index)
	} else {
		h.items[index] = val
		h.heapifyDown(index)
	}
}
