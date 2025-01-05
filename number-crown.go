package main

import "fmt"

func NumberCrown(n int) {
	for i := 0; i < n; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Print(j + 1)
		}

		for s := 0; s < 2*(n-i-1); s++ {
			fmt.Print(" ")
		}

		for k := i; k >= 0; k-- {
			fmt.Print(k + 1)
		}

		fmt.Println()
	}
}
