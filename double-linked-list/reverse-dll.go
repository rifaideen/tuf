package main

func reverseNode(head *Node) *Node {
	for {
		if head.Next == nil {
			head.Prev = head.Next
			head.Next = nil

			return head
		}

		head.Next, head.Prev = head.Prev, head.Next

		head = head.Prev
	}
}
