package main

func moveZeros(nums []int) {
	left := -1

	// find first 0th value and store it's index and break
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left = nums[i]
			break
		}
	}

	// when left is still -1, there are no zeros found
	if left == -1 {
		return
	}

	// using two pointer, left points to first 0th value index
	// right points to left + 1 until n
	// check if right pointer is pointing to non zero and swap and increment the left pointer
	for right := left + 1; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
