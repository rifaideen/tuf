package main

/**
 * Delete middle element of a linked list.
 * for odd length list, delete the middle element
 * for even length list, delete the second middle element
 */
func deleteMiddle(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	slow := head
	fast := head
	var prev *Node

	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	prev.next = slow.next

	return head
}
