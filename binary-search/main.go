package main

import "fmt"

func main() {
	fmt.Println("Binary search:", BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("Binary search (recursive):", RecursiveBinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("Floor and Ceil of 8:", FloorCeil([]int{3, 4, 4, 7, 8, 10}, 8))
	fmt.Println("Floor and Ceil of 2:", FloorCeil([]int{3, 4, 4, 7, 8, 10}, 2))
	fmt.Println("First and Last Occurence of 8 in 5,7,7,8,8,10 is ", FirstAndLastOccurence([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println("Count occurence of 3 in 2, 2 , 3 , 3 , 3 , 3 , 4 = ", CountOccurence([]int{2, 2, 3, 3, 3, 3, 4}, 3))
	fmt.Println("Count occurence of 2 in 1, 1, 2, 2, 2, 2, 2, 3 = ", CountOccurence([]int{1, 1, 2, 2, 2, 2, 2, 3}, 2))
}
