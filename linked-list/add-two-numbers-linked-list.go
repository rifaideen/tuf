package main

/**
 * Definition for singly-linked list.
 * type Node struct {
 *     data int
 *     next *Node
 * }
 */
func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	// create dummy node and point the reference to it
	node := &Node{
		data: -1,
	}
	current := node

	// create new heads to reference the l1 and l2
	head1 := l1
	head2 := l2

	// start with initial carry of 0
	carry := 0

	for head1 != nil || head2 != nil {
		// start the sum with carry
		sum := carry

		// add it to sum when head1 is not nil
		if head1 != nil {
			sum += head1.data
			head1 = head1.next
		}

		// add it to sum when head2 is not nil
		if head2 != nil {
			sum += head2.data
			head2 = head2.next
		}

		// create new node and attach it to the current node
		n := &Node{}

		// we need to carry the value
		if sum > 9 {
			// store the modulo of 10
			n.data = sum % 10
			// carry the division of 10
			carry = sum / 10
		} else { // sum is less than 10, no need to carry
			n.data = sum
			carry = 0
		}

		// link the new node and move
		current.next = n
		current = current.next
	}

	// check if we have to carry
	if carry > 0 {
		current.next = &Node{
			data: carry,
		}
	}

	// dummy node's next is the newly formed list
	return node.next
}
