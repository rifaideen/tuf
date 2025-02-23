package main

func MergeSortedArrays(nums1, nums2 []int) {
	left := 0
	right := 0

	for left < len(nums1) && right < len(nums2) {
		val1 := nums1[left]
		val2 := nums2[right]

		if val1 < val2 {
			nums1[left], nums2[right] = val2, val1
			left++
		} else {
			right++
		}
	}
}
