package main

func reverseKGroup(head *Node, k int) *Node {
	current := head
	var prev *Node

	for current != nil {
		kThNode := getKthNode(current, k)

		if kThNode == nil {
			if prev != nil {
				prev.next = current
			}

			break
		}

		next := kThNode.next
		kThNode.next = nil

		reverse(current)

		if current == head {
			head = kThNode
		} else {
			prev.next = kThNode
		}

		prev = current
		current = next
	}

	return head
}

func getKthNode(head *Node, k int) *Node {
	i := 1
	current := head

	for current != nil && i < k {
		current = current.next
		i++
	}

	return current
}

func reverse(head *Node) {
	current := head
	var prev *Node

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
}
