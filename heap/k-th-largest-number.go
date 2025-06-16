package main

import "container/heap"

// IntHeap follows min-heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Top() int {
	return h[0]
}

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	// create min heap and push k number of items
	h := &IntHeap{}
	heap.Init(h)

	for num := range k {
		heap.Push(h, nums[num])
	}

	// loop the remaining and push if current item is greater than top
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Top() {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	// we have removed the small items upto k, now top is the k-th largest number
	return h.Top()
}
