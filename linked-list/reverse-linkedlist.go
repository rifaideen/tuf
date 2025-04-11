package main

// using iterative approach
func reverseList(head *Node) *Node {
	current := head
	var prev *Node

	for current != nil {
		// store reference to next node before overwriting
		next := current.next

		// reverse the link
		current.next = prev
		// update the previous node
		prev = current

		// move to next node
		current = next
	}

	return prev
}

func reverseListRecursive(head *Node) *Node {
	// Base case:
	// If the linked list is empty or has only one node,
	// return the head as it is already reversed.
	if head == nil || head.next == nil {
		return head
	}

	// Recursive step:
	// Reverse the linked list starting
	// from the second node (head->next)
	newHead := reverseListRecursive(head.next)

	// save the reference to next node
	next := head.next
	// make the next of next node points to current head
	next.next = head
	// break the link from head to next to avoid cycles
	head.next = nil

	return newHead
}
