package main

import "fmt"

func InvertedNumberedRightPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("%d ", j+1)
		}

		fmt.Println()
	}
}
