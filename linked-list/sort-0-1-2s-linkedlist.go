package main

func sortOneTwoThree(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	zeroes := &Node{
		data: -1,
	}
	head0 := zeroes

	ones := &Node{
		data: -1,
	}
	head1 := ones

	twos := &Node{
		data: -1,
	}
	head2 := twos

	current := head

	for current != nil {
		switch current.data {
		case 0:
			head0.next = current
			head0 = head0.next
		case 1:
			head1.next = current
			head1 = head1.next
		case 2:
			head2.next = current
			head2 = head2.next
		}

		current = current.next
	}

	if ones.next != nil {
		head0.next = ones.next
	} else {
		head0.next = twos.next
	}

	head1.next = twos.next
	head2.next = nil

	return zeroes.next
}
