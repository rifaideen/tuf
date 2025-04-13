package main

func findPairs(head *Node, target int) [][]int {
	ans := [][]int{}
	left := head
	right := head

	for right.Next != nil {
		right = right.Next
	}

	for left.Data < right.Data {
		if left.Data+right.Data == target {
			ans = append(ans, []int{left.Data, right.Data})

			// target found, move both pointers
			left = left.Next
			right = right.Prev
		} else if left.Data+right.Data > target { // sum > target, move from right to left
			right = right.Prev
		} else { // sum < target, move from left to right
			left = left.Next
		}
	}

	return ans
}
