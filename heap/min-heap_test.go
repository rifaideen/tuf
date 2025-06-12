package main

import (
	"fmt"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	heap := NewMinHeap()
	heap.items = []int{-1, 10, 20, 15, 12, 40, 25, 18, 3, 1}
	heap.size = 9
	heap.heapify()
	fmt.Println(heap.items[1:])
}

/*
	     1
	 2      3
   4  5   6   7

   size / 2 (int division)
*/
