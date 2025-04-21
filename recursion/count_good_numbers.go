package main

// https://leetcode.com/problems/count-good-numbers/editorial/
func countGoodNumbers(n int64) int {
	// possible prime numbers in odd positions
	// 2,3,5,7

	// possible even numbers in even position
	// 0,2,4,6,8

	evenIndices := int((n + 1) / 2)
	oddIndices := int(n / 2)

	// modulo of 10 to the power 9 + 7 as mentioned in the problem statement
	mod := int(1e9 + 7)

	return (pow(5, evenIndices) * pow(4, oddIndices)) % mod
}

func pow(x, n int) int {
	// modulo of 10 to the power 9 + 7 as mentioned in the problem statement
	mod := int(1e9 + 7)

	// anything to the power 0 becomes 1, mathematics
	if n == 0 {
		return 1
	}

	// calculate the power by reducing n into half
	half := pow(x, n/2)

	// calculate the full by multiplying half * half.
	//
	// apply the formula of ans % mod
	ans := (half * half) % mod

	if n%2 == 0 {
		return ans
	}

	return (ans * x) % mod
}
