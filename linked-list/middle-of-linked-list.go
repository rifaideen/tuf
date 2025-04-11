package main

func middleNode(head *Node) *Node {
	// slow pointer will be moving 1 node per iteration
	slow := head
	// fast pointer will be moving 2 nodes per iteration
	fast := head

	// upon reaching nil on fast or fast.next, break the loop and next is the mid of the node
	for {
		if fast == nil || fast.next == nil {
			break
		}

		slow = slow.next
		fast = fast.next.next
	}

	return slow
}
