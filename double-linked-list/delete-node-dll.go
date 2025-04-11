package main

// 1 based index
func deleteNode(head *Node, x int) *Node {
	current := head

	if x == 1 {
		if current != nil && current.Next != nil {
			nextNode := current.Next
			nextNode.Prev = nil

			return nextNode
		} else {
			return nil // Only one node, and it's deleted
		}
	}

	i := 1
	for current != nil {
		if i == x {
			previous := current.Prev
			nextNode := current.Next

			if previous != nil {
				previous.Next = nextNode
			}
			if nextNode != nil {
				nextNode.Prev = previous
			}

			// If the deleted node was the head, update the head
			if current == head {
				return nextNode
			}
			return head
		}

		current = current.Next
		i++
	}

	return head // Position not found, return original head
}
