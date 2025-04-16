package main

type LinkedList struct {
	Val   int
	Next  *LinkedList
	Child *LinkedList
}

func flattenList(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}

	head2 := flattenList(head.Next)

	return flattenNodes(head, head2)
}

func flattenNodes(l1, l2 *LinkedList) *LinkedList {
	node := &LinkedList{
		Val: -1,
	}

	current := node

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Child = l1
			current = l1

			l1 = l1.Child
		} else {
			current.Child = l2
			current = l2

			l2 = l2.Child
		}
	}

	if l1 != nil {
		current.Child = l1
	}

	if l2 != nil {
		current.Child = l2
	}

	return node.Child
}
