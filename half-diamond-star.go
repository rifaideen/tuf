package main

import "fmt"

func HalfDiamondStar(n int) {
	var l, i, j int

	for i = -n; i <= n; i++ {
		if i < 0 {
			l = -i
		} else {
			l = i
		}

		for j = 0; j < n-l+1; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}
