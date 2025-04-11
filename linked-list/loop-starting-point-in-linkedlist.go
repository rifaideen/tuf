package main

/**
 * Given the head of a linked list, find the starting point of the loop and return it
 */
func loopStartingPosition(head *Node) *Node {
	slow := head
	fast := head
	isLoop := false

	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			isLoop = true
			break
		}
	}

	if !isLoop {
		return nil
	}

	// reset the slow and walk in the same speed for both slow and fast
	slow = head

	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return slow
}
