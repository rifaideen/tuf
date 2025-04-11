package main

/**
 * Remove the nth node from the end of the list and return its head.
 */
func removeNthFromEnd(head *Node, n int) *Node {
	length := 0
	current := head

	// calculate the length of l.list
	for current != nil {
		length++
		current = current.next
	}

	// calculate the node index to be removed
	removable := length - n

	// handle the edge case when removable is 0
	if removable == 0 {
		head = head.next
		return head
	}

	current = head
	i := 0

	for current != nil && current.next != nil {
		// check if next index equal removable, break the chain and return
		if i+1 == removable {
			current.next = current.next.next

			return head
		}

		i++
		current = current.next
	}

	return head
}
