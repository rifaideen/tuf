package main

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 > n2 {
		// we always to process the array with less number of items, hence calling the method with arrays swapped
		return FindMedianSortedArrays(nums2, nums1)
	}

	left := 0                          // I don't wanna take anything from nums1 to see if the symmetry is valid
	right := n1                        // I want to take the maximum of n1 numbers to form the symmetry
	required_left := (n1 + n2 + 1) / 2 // number of items on the left side required to form the symmetry
	n := n1 + n2

	for left <= right {
		mid1 := (left + right) / 2
		mid2 := required_left - mid1
		l1 := math.MinInt // set min int to handle comparision when l1 is not present when mid1 is 0
		l2 := math.MinInt // set min int to handle comparision when l2 is not present when mid2 is 0
		r1 := math.MaxInt // set max int to handle comparision when r1 is not present when mid1 is at last index
		r2 := math.MaxInt // set max int to handle comparision when l1 is not present when mid2 is at last index

		if mid1 < n1 { // check if mid1 is within n1
			r1 = nums1[mid1]
		}

		if mid2 < n2 { // check if mid2 is within n2
			r2 = nums2[mid2]
		}

		if mid1-1 >= 0 { // check if mid1 - 1 is not going negative index
			l1 = nums1[mid1-1]
		}

		if mid2-1 >= 0 { // check if mid2 - 1 is not going negative index
			l2 = nums2[mid2-1]
		}

		// check if we have valid symmetry
		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 {
				return float64(max(l1, l2))
			}

			return float64(max(l1, l2)+min(r1, r2)) / 2.0
		} else if l1 > r2 { // l1 is higher than r2, trim the right side
			right = mid1 - 1
		} else { //l2 is higer than r1, trim the left side
			left = mid1 + 1
		}
	}

	return 0
}
