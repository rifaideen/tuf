package main

import "fmt"

func main() {
	fmt.Println("Binary search:", BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("Binary search (recursive):", RecursiveBinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("Floor and Ceil of 8:", FloorCeil([]int{3, 4, 4, 7, 8, 10}, 8))
	fmt.Println("Floor and Ceil of 2:", FloorCeil([]int{3, 4, 4, 7, 8, 10}, 2))
}
