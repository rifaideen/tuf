package main

// Construct Double Linked List from array and return head
func constructDll(nums []int) *Node {
	if len(nums) == 0 {
		return &Node{}
	}

	head := &Node{
		Data: nums[0],
	}

	current := head

	for i := 1; i < len(nums); i++ {
		node := &Node{
			Data: nums[i],
		}
		node.Prev = current
		current.Next = node

		// replace current with next
		current = current.Next
	}

	return head
}
