package main

import "fmt"

func main() {
	fmt.Println("Binary search:", BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("Binary search (recursive):", RecursiveBinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
}
