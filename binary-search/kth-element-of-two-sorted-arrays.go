package main

import "math"

// using brute force solution
func KthElement(nums1, nums2 []int, k int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	ans := []int{}

	left := 0
	right := 0

	for left < n1 && right < n2 {
		if nums1[left] < nums2[right] {
			ans = append(ans, nums1[left])
			left++
		} else {
			ans = append(ans, nums2[right])
			right++
		}
	}

	for left < n1 {
		ans = append(ans, nums1[left])
		left++
	}

	for right < n2 {
		ans = append(ans, nums2[right])
		right++
	}

	return ans[k-1]
}

func KthElementBS(nums1, nums2 []int, k int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 > n2 {
		return KthElementBS(nums2, nums2, k)
	}

	// for explanation https://youtu.be/D1oDwWCq50g?si=5bsoL0leLm-9OqRm&t=486
	// at max we can pickup the k number of elements, if k = 2 and n1 is 5
	// taking 5 from nums1 does not make sence we can we want only k of 2 only
	left := max(0, k-n2)
	right := min(k, n1)

	for left <= right {
		mid1 := (left + right) / 2
		mid2 := k - mid1

		l1 := math.MinInt
		l2 := math.MinInt
		r1 := math.MaxInt
		r2 := math.MaxInt

		if mid1-1 >= 0 {
			l1 = nums1[mid1-1]
		}

		if mid2-1 >= 0 {
			l2 = nums2[mid2-1]
		}

		if mid1 < n1 {
			r1 = nums1[mid1]
		}

		if mid2 < n2 {
			r2 = nums2[mid2]
		}

		if l1 <= r2 && l2 <= r1 {
			return max(l1, l2)
		} else if l1 > r2 {
			right = mid1 - 1
		} else {
			left = mid1 + 1
		}
	}

	return -1
}
