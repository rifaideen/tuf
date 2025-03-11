package main

func SplitArray(nums []int, k int) int {
	left := 0
	right := 0

	for _, num := range nums {
		left = max(left, num)
		right += num
	}

	for left <= right {
		mid := (left + right) / 2

		if split(nums, mid) <= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func split(nums []int, val int) int {
	splitted := 1
	count := 0

	for _, num := range nums {
		if count+num > val {
			splitted++
			count = num
		} else {
			count += num
		}
	}

	return splitted
}
