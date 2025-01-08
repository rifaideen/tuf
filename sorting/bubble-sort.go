package main

func BubbleSort(nums []int) {
	for i := len(nums) - 1; i >= 1; i-- {
		// when it's not swapped at all after the iteration, the input is sorted already
		swapped := false

		for j := 0; j <= i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
