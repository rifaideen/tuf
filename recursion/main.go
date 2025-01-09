package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 3, "enter the value of n")

	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	fmt.Printf("Factorial(%d) = %d\n", *n, Factorial(*n))
	fmt.Printf("Sum of %d = %d\n", *n, SumOfN(*n))

	nums := []int{1, 2, 3, 4, 5}
	ReverseArray(nums, 0, len(nums)-1)

	fmt.Printf("Reverse Array = %v\n", nums)

	fmt.Printf("Check Palindrome(%q) = %v\n", "Hello", CheckPalindrome("Hello", 0))
	fmt.Printf("Check Palindrome(%q) = %v\n", "aha", CheckPalindrome("aha", 0))

	fmt.Printf("Fibonacci(%d) = %v\n", *n, Fibonacci(*n))
}

func Factorial(n int) int {
	if n < 0 {
		panic("n cannot be less than zero to perform factorial")
	}

	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func SumOfN(n int) int {
	if n <= 0 {
		return 0
	}

	return n + SumOfN(n-1)
}

func ReverseArray(nums []int, start, end int) {
	if start < end {
		// reverse in-place
		nums[start], nums[end] = nums[end], nums[start]

		// move the pointers
		ReverseArray(nums, start+1, end-1)
	}
}

func CheckPalindrome(str string, i int) bool {
	// if i exceeds half way, all items are compared
	if i >= len(str)/2 {
		return true
	}

	if str[i] != str[len(str)-1-i] {
		return false
	}

	return CheckPalindrome(str, i+1)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	last := Fibonacci(n - 1)
	secondLast := Fibonacci(n - 2)

	return last + secondLast
}
