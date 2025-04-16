package main

func rotateRight(head *Node, k int) *Node {
	if head == nil || head.next == nil || k == 0 {
		return head
	}

	current := head
	length := 1

	// calculate the length of linked list
	for current.next != nil {
		length++
		current = current.next
	}

	// attach the tail to the head to form the loop
	current.next = head

	// normalize the distance when the k is greather than length of linked list
	distance := k % length
	// calculate at what index we need to break the loop
	n := length - distance

	// move the current upto n times to find the new head
	for range n {
		current = current.next
	}

	// attach the new head and break the loop
	head = current.next
	current.next = nil

	// return the new head
	return head
}
