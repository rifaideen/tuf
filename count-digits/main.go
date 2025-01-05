package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CountDigits(12345))

	fmt.Println(CountDigits(7789))
}

func CountDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}
