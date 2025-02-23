package main

import (
	"fmt"
)

// using brute-force t.c = O(n2)
func CountInversions(nums []int) int {
	ans := 0

	for left := range nums {
		for right := left + 1; right < len(nums); right++ {
			if right < len(nums) && nums[left] > nums[right] {
				fmt.Printf("(%d, %d)\n", nums[left], nums[right])
				ans++
			}
		}
	}

	return ans
}

func CountInversionsOptimized(nums []int) int {
	res := mergesort(nums, 0, len(nums)-1)

	return res
}

func mergesort(nums []int, left, right int) int {
	count := 0

	if left >= right {
		return count
	}

	mid := (left + right) / 2

	count += mergesort(nums, left, mid)    // left half
	count += mergesort(nums, mid+1, right) // right half
	count += merge(nums, left, mid, right) // merge sorted halves

	return count
}

func merge(nums []int, low, mid, high int) int {
	count := 0
	ans := []int{}
	left := low
	right := mid + 1

	// while low < mid and mid + 1 > high
	for left <= mid && right <= high {
		if nums[left] <= nums[right] {
			ans = append(ans, nums[left])
			left++
		} else {
			ans = append(ans, nums[right])
			right++
			count += mid - left + 1
		}
	}

	// handle pointer exhaustion
	for left <= mid {
		ans = append(ans, nums[left])
		left++
	}

	for right <= high {
		ans = append(ans, nums[right])
		right++
	}

	for i := low; i <= high; i++ {
		nums[i] = ans[i-low]
	}

	return count
}
