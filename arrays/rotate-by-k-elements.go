package main

func rotateLeft(nums []int, k int) {
	n := len(nums)
	k = k % n

	// reverse first k elements
	reverse(nums, 0, k-1)
	// reverse the remaining elements
	reverse(nums, k, n-1)
	// reverse entire array
	reverse(nums, 0, n-1)
}

func rotateRight(nums []int, k int) {
	n := len(nums)
	k = k % n

	// reverse first group
	reverse(nums, 0, n-1-k)
	// reverse second group
	reverse(nums, n-k, n-1)
	// reverse entire array
	reverse(nums, 0, n-1)
}

func reverse(nums []int, start, end int) {
	// using two-pointer, upon each iteration these pointers move towards each other
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
