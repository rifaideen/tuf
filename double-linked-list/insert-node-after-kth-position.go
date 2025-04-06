package main

func insertAfter(head *Node, position, value int) {
	current := head
	i := 1

	for {
		if i-1 == position {
			node := &Node{
				Prev: current,
				Data: value,
				Next: current.Next,
			}
			current.Next = node
			break
		}

		i++
		current = current.Next
	}
}
