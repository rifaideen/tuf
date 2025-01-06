package main

import (
	"flag"
	"fmt"
	"math"
)

var n *int

func init() {
	n = flag.Int("n", 153, "enter the value of n")

	flag.Parse()
}

func main() {
	fmt.Printf("Is %d armstrong number? %v\n", *n, ArmstrongNumber(*n))
}

func ArmstrongNumber(n int) bool {
	sum := 0
	original := n
	pow := math.Log10(float64(n)) + 1

	for n > 0 {
		reminder := n % 10

		sum += int(math.Pow(float64(reminder), pow))

		n /= 10
	}

	return sum == original
}
