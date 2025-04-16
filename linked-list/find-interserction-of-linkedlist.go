package main

/**
 * Definition for singly-linked list.
 * type Node struct {
 * Val int
 * next *Node
 * }
 */
func getIntersectionNode(headA, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}

	len1, len2 := count(headA, headB)

	currentA := headA
	currentB := headB

	if len1 < len2 {
		// move len2
		currentB = move(currentB, len2-len1)
	} else if len1 > len2 {
		// move len1
		currentA = move(currentA, len1-len2)
	}

	for currentA != nil && currentB != nil {
		if currentA == currentB {
			return currentA
		}

		currentA = currentA.next
		currentB = currentB.next
	}

	return nil
}

// count the length of given linked lists
func count(headA, headB *Node) (int, int) {
	countA := 0
	countB := 0

	// count both of them at the same time
	for headA != nil || headB != nil {
		if headA != nil {
			countA++
			headA = headA.next
		}

		if headB != nil {
			countB++
			headB = headB.next
		}
	}

	return countA, countB
}

func move(head *Node, position int) *Node {
	for i := 0; i < position; i++ {
		head = head.next
	}

	return head
}
