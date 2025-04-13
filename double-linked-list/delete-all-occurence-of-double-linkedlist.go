package main

func deleteAllOccurences(head *Node, target int) *Node {
	// make a reference to the head
	current := head

	for current != nil {
		// check if data equals to target
		if current.Data == target {
			// move head when current and head are same
			if current == head {
				head = head.Next
			}

			prev := current.Prev
			next := current.Next

			// conditionally update the links
			if next != nil {
				next.Prev = prev
			}

			if prev != nil {
				prev.Next = next
			}
		}

		// proceed further
		current = current.Next
	}

	return head
}
