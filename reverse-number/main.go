package main

import "fmt"

func main() {
	fmt.Println(ReverseNumbers(1234567890))
}

func ReverseNumbers(n int) int {
	reversed := 0

	for n > 0 {
		// find the modulo division reminder
		reminder := n % 10

		// multiply the reversed value by 10 and add reminder
		reversed = reversed*10 + reminder

		// update n by n / 10
		n = n / 10
	}

	return reversed
}
