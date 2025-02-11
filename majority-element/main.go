package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("MajorityElement(%v) = %v\n", nums, MajorityElement(nums))
	fmt.Printf("MajorityElementII(%v) = %v\n", nums, MajorityElementII(nums))
}

// booyer-moore algorithm
func MajorityElement(nums []int) int {
	count := 0
	element := -math.MaxInt

	for i := range nums {
		if count == 0 {
			element = nums[i]
			count++
		} else if element == nums[i] {
			count++
		} else {
			count--
		}
	}

	return element
}

func MajorityElementII(nums []int) []int {
	// create 2 counters and 2 elements with default values
	el1 := math.MaxInt // using MaxInt instead of None or 0 because 0 maybe part of nums
	cnt1 := 0

	el2 := math.MaxInt // using MaxInt instead of None or 0 because 0 maybe part of nums
	cnt2 := 0

	// create map to store the frequency of each number in nums to be used later
	frequency := map[int]int{}

	for _, num := range nums {
		frequency[num]++

		// check if counter 1 is zero and current num is not same as el2
		if cnt1 == 0 && el2 != num {
			cnt1++
			el1 = num
		} else if cnt2 == 0 && el1 != num { // check if counter 2 is zero and current num is not same as el1
			cnt2++
			el2 = num
		} else if el1 == num { // check if num is el1
			cnt1++
		} else if el2 == num { // check if nums is el2
			cnt2++
		} else { // none of them match, decrement the counters
			cnt1--
			cnt2--
		}
	}

	res := []int{}

	// check the frequency of el1 > n/3
	if frequency[el1] > len(nums)/3 {
		res = append(res, el1)
	}

	// check the frequency of el2 > n/3
	if frequency[el2] > len(nums)/3 {
		res = append(res, el2)
	}

	return res
}
