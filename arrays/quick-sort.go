package main

func QuickSort(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, start, end int) {
	if start < end {
		// create partition index
		i := partition(nums, start, end)

		// quick sort from start until i -1
		quicksort(nums, start, i-1)
		// quick sort from i+1 to end
		quicksort(nums, i+1, end)
	}
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]

	// two-pointers
	left, right := start, end

	for left < right {
		// move the left pointer
		for nums[left] <= pivot && left < end {
			left++
		}

		// move the right pointer
		for nums[right] > pivot && right > start {
			right--
		}

		// still left is less than right? swap here.
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	// swap start with right
	nums[start], nums[right] = nums[right], nums[start]

	return right
}
