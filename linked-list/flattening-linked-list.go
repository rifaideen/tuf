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
			current.Child = &LinkedList{
				Val: l1.Val,
			}
			current = current.Child
			l1 = l1.Child
		} else {
			current.Child = &LinkedList{
				Val: l2.Val,
			}
			current = current.Child

			l2 = l2.Child
		}
	}

	for l1 != nil {
		current.Child = &LinkedList{
			Val: l1.Val,
		}
		current = current.Child
		l1 = l1.Child
	}

	for l2 != nil {
		current.Child = &LinkedList{
			Val: l2.Val,
		}
		current = current.Child

		l2 = l2.Child
	}

	return node.Child
}
