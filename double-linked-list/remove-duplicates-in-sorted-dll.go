package main

func removeDuplicates(head *Node) *Node {
	current := head

	for current != nil && current.Next != nil {
		next := current.Next

		for next != nil && current.Data == next.Data {
			next = next.Next
		}

		current.Next = next

		if next != nil {
			next.Prev = current
		}

		current = current.Next
	}

	return head
}
