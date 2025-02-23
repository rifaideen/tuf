package main

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// divide the array
	l, r := divide(nums)

	// process the l and r
	left := MergeSort(l)
	right := MergeSort(r)

	// merge left and right
	return conquer(left, right)
}

func divide(nums []int) (left, right []int) {
	mid := len(nums) / 2

	for i := range nums {
		if i < mid {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	return
}

// conquer uses two-pointer algorithm to merge the two arrays in sorted fashion
func conquer(left, right []int) []int {
	t := []int{}
	i := 0
	j := 0

	// when looping two arrays using two-pointers, we might miss out one of the array
	// when left and right array sizes are not same
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			t = append(t, left[i])
			i++
		} else {
			t = append(t, right[j])
			j++
		}
	}

	// cover the exhaused left pointer
	for i < len(left) {
		t = append(t, left[i])
		i++
	}

	// cover the exhaused right pointer
	for j < len(right) {
		t = append(t, right[j])
		j++
	}

	return t
}
