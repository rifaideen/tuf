package main

func hasCycle(head *Node) bool {
	slow := head
	fast := head

	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}
