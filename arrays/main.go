package main

import "fmt"

func main() {
	nums0 := []int{3, 4, 5, 1, 2}
	fmt.Println("Is it sorted? ", check(nums0))

	removeDuplicates([]int{1, 1, 2})

	nums := []int{1, 2, 3, 4, 5}
	rotateOneLeft(nums)

	fmt.Println("after rotating one element to left", nums)

	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	rotateLeft(nums2, 3)
	fmt.Println("after rotating 3 elements to left", nums2)

	nums3 := []int{1, 2, 3, 4, 5, 6, 7}
	rotateRight(nums3, 3)
	fmt.Println("after rotating 3 elements to right", nums3)

	nums4 := []int{0, 1, 0, 3, 12}
	moveZeros(nums4)
	fmt.Println("after moving zeroes to end", nums4)

	nums5 := []int{7, 5, 4, 2, 1, 8}
	QuickSort(nums5)
	fmt.Println("after quick sort", nums5)

	nums6 := []int{1, 2, 3, 4, 6}
	nums7 := []int{2, 3, 5}
	fmt.Printf("union of %v and %v is %v\n", nums6, nums7, findUnion(nums6, nums7))

	missingSortedNumber([]int{1, 2, 4, 5}, 5)

	missingNumber([]int{1, 2, 4, 5})

	nums = []int{1, 1, 2, 3, 1, 1, 1, 1, 4, 5, 6, 1, 1}
	fmt.Printf(
		"Max consecutive ones in %v is %d\n",
		nums,
		findMaxConsecutiveOnes(nums),
	)

	// N = 3, k = 5, array[] = {2,3,5}
	fmt.Println("longest sub-array", longestSubarray([]int{-1, 1, 1}, 1))

	arr := [][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}

	Intersection(arr)

	LargestUniqueNumber([]int{11, 10, 11})

	MaxBaloons("loonbalxballpoon")

	LongestConsecutive([]int{100, 4, 200, 1, 3, 2})

	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	MatrixZero(matrix)
	fmt.Println("After matrix zero", matrix)

	spiral := SpiralMatrix([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})

	fmt.Println("Sprial matrix value", spiral)

	fmt.Println("Get pascal triangle of 5:")

	for _, values := range GetPascalTriangle(5) {
		for _, val := range values {
			fmt.Print(val, " ")
		}

		fmt.Println()
	}

	fmt.Println("Get pascal element at 5th row, 3rd col:", GetPascalElement(5, 3))
	fmt.Println("Get pascal row for 5th row", GetPascalRow(4))

	nums = []int{9, -3, 3, -1, 6, -5}
	// nums = []int{6, -2, 2, -8, 1, 7, 4, -10}
	fmt.Println("LongestSubArraySumZero(nums):", LongestSubArraySumZero(nums))
}

// Given an array nums, return true if the array was originally sorted in non-decreasing order, then rotated some number of positions (including zero). Otherwise, return false.
//
// There may be duplicates in the original array.
//
// Note: An array A rotated by x positions results in an array B of the same length such that A[i] == B[(i+x) % A.length], where % is the modulo operation.
func check(nums []int) bool {
	inversions := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			inversions++
		}

		if inversions > 1 {
			return false
		}
	}

	return true
}

func removeDuplicates(nums []int) int {
	left := 0

	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left += 1
			nums[left] = nums[right]
		}
	}

	return left + 1
}
