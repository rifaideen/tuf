package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
// The idea is to create a min heap which stores the list of ListNode and upon heapifying up or down
// we will take advantage of min heap's top being smallest.
//
// we compare the current node's top value and parent node's top value, similary for left and right node's value.
//
// # We start the loop and insert the list node one by one and min heap to keep them in order
//
// The legendary part is, while merging the nodes back, we pop the smallest element and assign it to the merged node
// and if the popped node has next elements, we put that next value back to the heap.
//
// This way, heap has the opportunity to explore all the values of the list, not just the top, left or right.
func mergeKLists(lists []*ListNode) *ListNode {
	// handle egde case
	if len(lists) == 0 {
		return nil
	}

	h := &MinHeap{
		size: 0,
		nums: make([]*ListNode, 0),
	}

	for _, list := range lists {
		// push non empty list to the heap
		if list != nil {
			h.push(list)
		}
	}

	// immediately stop when the heap is empty
	if h.empty() {
		return nil
	}

	// create dummy node and assign it to current
	head := &ListNode{}
	current := head

	// pop the values one by one and push the next nodes back to the heap
	for !h.empty() {
		// keep a reference of popped node
		min := h.pop()

		// push the next value back to the heap
		if min.Next != nil {
			h.push(min.Next)
		}

		// assign the popped node and move current node
		current.Next = min
		current = current.Next
	}

	return head.Next
}

// MinHeap represent the min-heap data structure with list of ListNode
type MinHeap struct {
	size int
	nums []*ListNode
}

func (h *MinHeap) push(num *ListNode) {
	h.nums = append(h.nums, num)
	h.up(h.size)
	h.size++
}

func (h *MinHeap) pop() *ListNode {
	if h.size == 0 {
		panic("heap empty")
	}

	num := h.nums[0]

	h.size--
	h.nums[0] = h.nums[h.size]
	h.nums = h.nums[:h.size]

	h.down(0)

	return num
}

func (h *MinHeap) up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		// we take advantage of min-heap data structure of top being the smallest
		// compare current and parent value
		if h.nums[index].Val < h.nums[parent].Val {
			h.swap(index, parent)
			h.up(parent)
		}
	}
}

func (h *MinHeap) down(index int) {
	left := 2*index + 1
	right := left + 1
	smallest := index

	// compare current and left node value
	if left < h.size && h.nums[left].Val < h.nums[smallest].Val {
		smallest = left
	}

	// compare current and right node value
	if right < h.size && h.nums[right].Val < h.nums[smallest].Val {
		smallest = right
	}

	if index != smallest {
		h.swap(index, smallest)
		h.down(smallest)
	}
}

func (h *MinHeap) swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *MinHeap) empty() bool {
	return h.size == 0
}
