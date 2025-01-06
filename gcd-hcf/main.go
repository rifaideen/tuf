package main

import "fmt"

func GCD(num1, num2 int) int {
	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			num1 = num1 % num2
		} else {
			num2 = num2 % num1
		}
	}

	if num1 == 0 {
		return num2
	}

	return num1
}

func main() {
	fmt.Println("GCD(9,12) = ", GCD(9, 12))
	fmt.Println("GCD(20,15) = ", GCD(20, 15))
}
