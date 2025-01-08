package main

import "fmt"

func main() {
	nums := []int{54, 97, 674, 8, 6, 4}
	BubbleSort(nums)
	fmt.Println("After Bubble Sorting = ", nums)

	nums2 := []int{4, 6, 8, 54, 97, 674}
	BubbleSort(nums2)
	fmt.Println("After Bubble Sorting = ", nums2)
}
