package main

func constructLL(nums []int) *Node {
	head := &Node{
		data: nums[0],
	}

	current := head

	for i := 1; i < len(nums); i++ {
		node := &Node{
			data: nums[i],
		}

		current.next = node
		current = node
	}

	return head
}
