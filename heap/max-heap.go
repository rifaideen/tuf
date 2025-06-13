package main

type MaxHeap struct {
	items []int
	size  int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		items: []int{},
		size:  0,
	}
}

func NewMaxHeapFrom(items []int) *MaxHeap {
	h := NewMaxHeap()
	h.items = items
	h.size = len(items)
	h.heapify()

	return h
}

func (h *MaxHeap) heapifyUp(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.items[index] > h.items[parent] {
			h.items[index], h.items[parent] = h.items[parent], h.items[index]
			h.heapifyUp(parent)
		}
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	left := 2*index + 1
	right := left + 1
	smallest := index

	if left < h.size && h.items[left] > h.items[smallest] {
		smallest = left
	}

	if right < h.size && h.items[right] > h.items[smallest] {
		smallest = right
	}

	if smallest != index {
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]

		h.heapifyDown(smallest)
	}
}

func (h *MaxHeap) heapify() {
	for i := h.size / 2; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func (h *MaxHeap) Insert(key int) {
	h.items = append(h.items, key)
	h.heapifyUp(h.size)
	h.size++
}

func (h *MaxHeap) GetMax() int {
	if h.size == 0 {
		return -1
	}

	return h.items[0]
}

// delete and return
func (h *MaxHeap) ExtractMax() int {
	if h.size == 0 {
		return -1
	}

	top := h.items[0]
	h.size--
	h.items[0] = h.items[h.size]
	h.items = h.items[:h.size]

	h.heapifyDown(0)

	return top
}

// size of heap
func (h *MaxHeap) HeapSize() int {
	return h.size
}

func (h *MaxHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MaxHeap) ChangeKey(index, val int) {
	current := h.items[index]
	h.items[index] = val

	if val > current {
		h.heapifyUp(index)
	} else {
		h.heapifyDown(index)
	}
}
