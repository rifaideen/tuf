package main

import (
	"fmt"
	"strings"
)

// Generate all binary strings without consecutive 1’s
// Start with n = 1
// if n == k ? print str
// if str[n-1] == 1 ? str[n] = 0 and re-generate with n+1
// if str[n-1] == 0 ? str[n] = 1 and re-generate with n+1 and str[n] = 0 and re-generate with n+1

// This approach has high space complexity, which is not accepted
// T.C O(2^k)
// S.C O(K)
func generateBinaryStrings(k int) {
	if k <= 0 {
		return
	}

	// start with 1
	n := 1
	str := make([]string, k)

	// generate the binary strings for 0
	str[0] = "0"
	generateKBinaryStrings(k, n, str)

	// generate the binary strings for 1
	str[0] = "1"
	generateKBinaryStrings(k, n, str)
}

func generateKBinaryStrings(k, n int, str []string) {
	// when n equals k, print the strings and add space at the end
	if n == k {
		fmt.Print(strings.Join(str, "") + " ")
		return
	}

	// when last character is 1, we can add only 0 and generate the binary strings
	if str[n-1] == "1" {
		str[n] = "0"
		generateKBinaryStrings(k, n+1, str)
		return
	}

	// when last character is 0, we should generate for 1 and 0 recursively
	if str[n-1] == "0" {
		str[n] = "1"
		generateKBinaryStrings(k, n+1, str)

		str[n] = "0"
		generateKBinaryStrings(k, n+1, str)
	}
}

// Optimized approach
// T.C O(2^K)
// S.C O(N)
func generateBinaryStringsOptimized(k int) {
	if k <= 0 {
		return
	}

	nums := make([]bool, k)

	n := 1

	nums[0] = false
	generateKBinaryStringsOptimized(k, n, nums)

	nums[0] = true
	generateKBinaryStringsOptimized(k, n, nums)
}

func generateKBinaryStringsOptimized(k, n int, nums []bool) {
	if n == k {
		for _, v := range nums {
			if v {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}

		fmt.Print(" ")
		return
	}

	if nums[n-1] {
		nums[n] = false
		generateKBinaryStringsOptimized(k, n+1, nums)

		return
	}

	if !nums[n-1] {
		nums[n] = false
		generateKBinaryStringsOptimized(k, n+1, nums)

		nums[n] = true
		generateKBinaryStringsOptimized(k, n+1, nums)
	}
}
