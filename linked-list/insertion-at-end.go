package main

func insertEnd(head *Node, num int) *Node {
	current := head

	for current != nil {
		if current.next != nil {
			current = current.next
		} else {
			node := &Node{
				data: num,
			}

			current.next = node
			current = current.next
		}
	}

	return head
}
