package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0, 2}
	nums2 := []int{1, 0, 2}
	nums3 := []int{1, 0, 2}

	sortColorsTwoPointer(nums)
	fmt.Println("sorted two-pointer:", nums)

	sortColorsBrute(nums2)
	fmt.Println("sorted brute-force:", nums2)

	sortOptimal(nums3)
	fmt.Println("sorted optimal:", nums3)
}

func sortColorsTwoPointer(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		if nums[mid] == 0 { // for 0s
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 { // for 1s
			mid++
		} else { // for 2s
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func sortColorsBrute(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func sortOptimal(nums []int) {
	zeroes, ones, twos := 0, 0, 0

	// count the number of 0,1,2 and arrange in ascending order
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			zeroes++
		case 1:
			ones++
		default:
			twos++
		}
	}

	i := 0

	for ; i < zeroes; i++ {
		nums[i] = 0
	}

	for ; i < zeroes+ones; i++ {
		nums[i] = 1
	}

	for ; i < zeroes+ones+twos; i++ {
		nums[i] = 2
	}
}
