package main

import "fmt"

func DiamondStarCombined(n int) {
	Pyramid(n)
	InvertedStarPyramid(n)
}

func DiamondStar(n int) {
	var l int

	for i := -n; i <= n; i++ {
		if i < 0 {
			l = -i
		} else {
			l = i
		}

		for s := 0; s < l; s++ {
			fmt.Print("  ")
		}

		for j := 0; j < (2*(n-l))+1; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}

func DiamondNumberedIncremental(n int) {
	var l int

	for i := -n; i <= n; i++ {
		if i < 0 {
			l = -i
		} else {
			l = i
		}

		for s := 0; s < l; s++ {
			fmt.Print("  ")
		}

		for j := 0; j < (2*(n-l))+1; j++ {
			fmt.Printf("%d ", j+1)
		}

		fmt.Println()
	}
}

func DiamondNumberedConstant(n int) {
	var l int

	for i := -n; i <= n; i++ {
		if i < 0 {
			l = -i
		} else {
			l = i
		}

		for s := 0; s < l; s++ {
			fmt.Print("  ")
		}

		for j := 0; j < (2*(n-l))+1; j++ {
			fmt.Printf("%d ", l)
		}

		fmt.Println()
	}
}
