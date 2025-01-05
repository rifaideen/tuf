package main

import "fmt"

func Pyramid(n int) {
	for i := 0; i < n; i++ {
		// print the spaces for each row
		for s := 0; s < n-i-1; s++ {
			fmt.Print("  ")
		}

		for j := 0; j < (i*2)+1; j++ {
			fmt.Print("* ")
		}

		// move to new line
		fmt.Println()
	}
}
