package main

import "fmt"

func RightAngledNumberPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", j+1)
		}

		fmt.Println()
	}
}

func RightAngledNumberPyramid2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", i+1)
		}

		fmt.Println()
	}
}
