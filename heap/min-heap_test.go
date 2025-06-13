package main

import (
	"fmt"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	heap := NewMinHeap()
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(10)

	fmt.Println(heap.items)
	fmt.Println(heap.GetMin())
	fmt.Println(heap.HeapSize())
	fmt.Println(heap.IsEmpty())
	fmt.Println(heap.ExtractMin())
	heap.ChangeKey(1, 16)
	fmt.Println(heap.items)
	fmt.Println(heap.GetMin())
	fmt.Println(heap.items)

	h := NewMinHeapFrom([]int{4, 16, 10, 1, 12, 16})
	fmt.Println(h.items)
}
