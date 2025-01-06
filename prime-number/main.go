package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Is 1483 prime number? (brute-force) ", IsPrimeNumber(1483))
	fmt.Println("Is 1483 prime number? (optimal) ", IsPrimeNumberOptimized(1483))
}

func IsPrimeNumber(n int) bool {
	count := 0

	for i := 1; i <= n; i++ {
		// we don't want to process further when count is more than 2
		if count > 2 {
			return false
		}

		if n%i == 0 {
			count++
		}
	}

	return true
}

func IsPrimeNumberOptimized(n int) bool {
	count := 0
	sqrt := math.Sqrt(float64(n))

	for i := 1; i <= int(sqrt); i++ {
		if count > 2 {
			return false
		}

		if n%i == 0 {
			count++

			if n/i != i {
				count++
			}
		}
	}

	return true
}
