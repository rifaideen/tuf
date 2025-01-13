package main

func rotateOneLeft(nums []int) {
	n := len(nums)

	if n <= 1 {
		return
	}

	t := nums[0]

	for i := 1; i < n; i++ {
		nums[i-1] = nums[i]
	}

	nums[n-1] = t
}
