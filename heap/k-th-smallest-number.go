package main

import "container/heap"

// IntHeapMax follows max-heap
type IntHeapMax []int

func (h IntHeapMax) Len() int           { return len(h) }
func (h IntHeapMax) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeapMax) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeapMax) Top() int {
	return h[0]
}

func (h *IntHeapMax) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeapMax) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findSmallest(nums []int, k int) int {
	h := &IntHeapMax{}
	heap.Init(h)

	for i := range k {
		heap.Push(h, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < h.Top() {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	return h.Top()
}
