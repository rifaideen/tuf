package main

// give head of a linkedlist, task is to add 1 to the overall linkedlist
//
// using recursive approach to solve the problem from the backward position
func addOne(head *Node) *Node {
	// add one to the node recursively
	carry := addOneToNode(head)

	// if the final carry is one, we need to add new node and point the next node to the head
	if carry == 1 {
		return &Node{
			data: carry,
			next: head,
		}
	}

	return head
}

// recursively add one to the node and return the carry
func addOneToNode(head *Node) int {
	// upon reaching the end of linkedlist, we send 1 as carry
	if head == nil {
		return 1
	}

	carry := addOneToNode(head.next)

	head.data += carry

	// check if data + carry equals to 10
	if head.data == 10 {
		// add zero to the data, and return 1 to carry
		head.data = 0

		return 1
	}

	// data + carry is less than 10, add it and return 0 to carry
	return 0
}
