package main

import "fmt"

func main() {
	fmt.Println("Is 121 palindrome? ", IsPalindrome(121))
	fmt.Println("Is 321 palindrome? ", IsPalindrome(321))
	fmt.Println("Is 42024 palindrome? ", IsPalindrome(42024))
}

// Reverse the given number and check with original to see if the given number is palindrome or not
func IsPalindrome(n int) bool {
	reversed := 0
	original := n

	for n > 0 {
		lastDigit := n % 10

		reversed = reversed*10 + lastDigit

		n = n / 10
	}

	return reversed == original
}
