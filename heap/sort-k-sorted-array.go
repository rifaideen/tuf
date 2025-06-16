package main

func nearlySorted(nums []int, k int) {
	h := NewMinHeap()

	// insert k number of items into the heap
	for i := range k {
		h.Insert(nums[i])
	}

	var i int

	// loop the remaining from k
	for i = k; i < len(nums); i++ {
		// insert the item, since heap contains k+1 items, remove the min item
		h.Insert(nums[i])
		nums[i-k] = h.ExtractMin()
	}

	// remove the remaining items
	for !h.IsEmpty() {
		nums[i-k] = h.ExtractMin()
		i++
	}
}
