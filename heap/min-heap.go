package main

type MinHeap struct {
	items []int
	size  int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: []int{},
		size:  0,
	}
}

func NewMinHeapFrom(items []int) *MinHeap {
	h := NewMinHeap()
	h.items = items
	h.size = len(items)
	h.heapify()

	return h
}

func (h *MinHeap) heapifyUp(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.items[index] < h.items[parent] {
			h.items[index], h.items[parent] = h.items[parent], h.items[index]
			h.heapifyUp(parent)
		}
	}
}

func (h *MinHeap) heapifyDown(index int) {
	left := 2*index + 1
	right := left + 1
	smallest := index

	if left < h.size && h.items[left] < h.items[smallest] {
		smallest = left
	}

	if right < h.size && h.items[right] < h.items[smallest] {
		smallest = right
	}

	if smallest != index {
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]

		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) heapify() {
	// scanning from right to left. since leaf nodes won't have any child to compare downwards
	// we use h.size/2 to start immediately above the leaf nodes
	for i := h.size / 2; i >= 0; i-- {
		// this will compare the current index with left and right children
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

	// take a backup of top element before removing it
	top := h.items[0]
	// reduce the size
	h.size--

	// put the last item into the root and remove the last element
	h.items[0] = h.items[h.size]
	h.items = h.items[:h.size]

	// heapify from top to bottom
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
	current := h.items[index]
	h.items[index] = val

	if val < current {
		h.heapifyUp(index)
	} else {
		h.heapifyDown(index)
	}
}
