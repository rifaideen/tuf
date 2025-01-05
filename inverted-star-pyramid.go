package main

import "fmt"

func InvertedStarPyramid(n int) {
	// 1 + (4 * 2)
	for i := 0; i < n; i++ {
		// print spaces
		for s := 0; s < i; s++ {
			fmt.Print("  ")
		}

		// print stars
		for j := 0; j < 2*(n-i)-1; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
