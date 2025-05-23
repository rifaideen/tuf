package main

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		k := nums[i]
		var j int

		for j = i - 1; j >= 0 && nums[j] > k; j-- {
			nums[j+1] = nums[j]
		}

		nums[j+1] = k
	}
}

func RecursiveInsertionSort(nums []int, start, n int) {
	if start == n {
		return
	}

	j := start

	for j > 0 && nums[j-1] > nums[j] {
		nums[j], nums[j-1] = nums[j-1], nums[j]
		j--
	}

	RecursiveInsertionSort(nums, start+1, n)
}
