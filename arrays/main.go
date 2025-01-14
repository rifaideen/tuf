package main

import "fmt"

func main() {
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
}
