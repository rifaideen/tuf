package main

// 54, 97, 674, 8, 6, 4
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		// wheter the input is sorted already
		swapped := false

		for j := 0; j < (len(nums)-1)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}

		// no swap happened
		if !swapped { // input is sorted already, break the loop
			break
		}
	}
}
