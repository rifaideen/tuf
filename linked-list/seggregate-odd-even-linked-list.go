package main

// seggregate odd and even nodes in a linked list
func oddEvenList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	odd := head
	even := head.next
	even_head := head.next

	for even != nil && even.next != nil {
		odd.next = odd.next.next
		even.next = even.next.next

		odd = odd.next
		even = even.next
	}

	odd.next = even_head

	return head
}
