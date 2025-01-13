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
}
