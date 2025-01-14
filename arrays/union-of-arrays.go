package main

func findUnion(nums1, nums2 []int) []int {
	sorted := []int{}
	left, right := 0, 0

	for left < len(nums1) && right < len(nums2) {
		if nums1[left] <= nums2[right] {
			if len(sorted) == 0 || sorted[len(sorted)-1] != nums1[left] {
				sorted = append(sorted, nums1[left])
			}

			left++
		} else {
			if len(sorted) == 0 || sorted[len(sorted)-1] != nums2[right] {
				sorted = append(sorted, nums2[right])
			}

			right++
		}
	}

	for left < len(nums1) {
		if len(sorted) == 0 || sorted[len(sorted)-1] != nums1[left] {
			sorted = append(sorted, nums1[left])
		}

		left++
	}

	for right < len(nums2) {
		if len(sorted) == 0 || sorted[len(sorted)-1] != nums2[right] {
			sorted = append(sorted, nums2[right])
		}

		right++
	}

	return sorted
}
