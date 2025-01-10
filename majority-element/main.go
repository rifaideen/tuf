package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 1, 1, 1, 2, 3, 5}
	fmt.Printf("MajorityElement(%v) = %v\n", nums, MajorityElement(nums))
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
