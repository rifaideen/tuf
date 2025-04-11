package main

// Sorting linked list using merge sort
func sortList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	// find the middle node
	middle := findMiddle(head)
	// next of middle node is our right node
	right := middle.next
	// unlink the next node from the middle to remove the chain
	middle.next = nil

	// now our head is splitted upto middle node (we set middle.next to nil so it won't go further)
	left := head

	left = sortList(left)
	right = sortList(right)

	return merge(left, right)
}

// using rabbit and hare algorithm to find the middle of the linked list
func findMiddle(head *Node) *Node {
	slow := head
	fast := head.next // since we always want to mid to point at the first mid (in case of even number of nodes), we had to start from head.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func merge(left, right *Node) *Node {
	nodes := &Node{}
	current := nodes

	for left != nil && right != nil {
		if left.data < right.data {
			current.next = left
			left = left.next
		} else {
			current.next = right
			right = right.next
		}

		current = current.next
	}

	for left != nil {
		current.next = left

		left = left.next
		current = current.next
	}

	for right != nil {
		current.next = right

		right = right.next
		current = current.next
	}

	return nodes.next
}
