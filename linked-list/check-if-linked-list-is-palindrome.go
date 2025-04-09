package main

/**
 * Brute-force solution using stack
 *
 * Push data into the stack and compare with head.
 * T.C O(2N)
 * S.C O(N)
 */
func isPalindrome(head *Node) bool {
	top := -1
	stack := []int{}

	// push them onto the stack
	current := head

	for current != nil {
		// push it to stack and update the top
		stack = append(stack, current.data)
		top++

		// move on to next node
		current = current.next
	}

	current = head

	for current != nil {
		if stack[top] != current.data {
			return false
		}

		current = current.next
		top--
	}

	return true
}

// reverse linked list, recursive
func reverseLinkedList(head *Node) *Node {
	// when head or head.next is nil, return immediately
	if head == nil || head.next == nil {
		return head
	}

	// reverse the next pointer
	newHead := reverseLinkedList(head.next)
	// store the head.next reference, before overwriting
	next := head.next

	next.next = head
	head.next = nil

	return newHead
}

func isPalindromeOptimized(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	// find the middle of linked list using tortoise & hare algorithm
	slow, fast := head, head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// reverse the next from slow pointer which would be the second half
	newHead := reverseLinkedList(slow.next)

	// compare the first and second half
	first := head
	second := newHead

	for second != nil {
		// data does not match, not a palindrome
		if first.data != second.data {
			reverseLinkedList(newHead)
			return false
		}

		first = first.next
		second = second.next
	}

	// reverse the reversed to make the original order
	reverseLinkedList(newHead)

	return true
}
