package main

import "fmt"

func main() {
	nums := []int{13, 46, 24, 52, 20, 9, 1}

	fmt.Printf("SelectionSort(%v)\n", nums)
	SelectionSort(nums)
	fmt.Println(nums)

	nums1 := []int{54, 97, 674, 8, 6, 4}

	fmt.Printf("BubbleSort(%v)\n", nums1)
	BubbleSort(nums1)
	fmt.Println(nums1)

	nums2 := []int{5, 8, 2, 4, 3, 78, 2, 1}

	fmt.Printf("BubbleSort(%v)\n", nums2)
	BubbleSort(nums2)
	fmt.Println(nums2)

	nums3 := []int{13, 46, 24, 52, 20, 9}

	fmt.Printf("InsertionSort(%v)\n", nums3)
	InsertionSort(nums3)
	fmt.Println(nums3)
}
