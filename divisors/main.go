package main

import "fmt"

func main() {
	fmt.Println("Divisors of 36 (brut force)", Divisor(36))
	fmt.Println("Divisors of 36 (brut force)", Divisor(36))

	fmt.Println()

	fmt.Println("Divisors of (optimal) 36", DivisorOptimal(36))
	fmt.Println("Divisors of (optimal) 12", DivisorOptimal(12))
}

// using brute-force
func Divisor(n int) []int {
	divisors := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

func DivisorOptimal(n int) []int {
	divisors := []int{}

	for i := 1; i*i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)

			// check if n/i != i and append n/i
			if (n / i) != i {
				divisors = append(divisors, n/i)
			}

		}
	}

	return divisors
}
