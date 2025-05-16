package main

import (
	"slices"
)

func primeFactors(nums []int) [][]int {
	ans := [][]int{}

	// find the max value in the nums and create new array of size max value + 1
	maxValue := slices.Max(nums)
	primes := make([]int, maxValue+1)

	// fill primes array with incremental values i.e 0,1,2,3...maxValue + 1
	// we do this to find the smallest prime factors of each numbers
	for value := range maxValue + 1 {
		primes[value] = value
	}

	// apply seive algorithm to find the smallest prime factors
	for i := 2; i*i <= maxValue+1; i++ { // loop from 2 to square root of maxValue+1
		for j := i * i; j < maxValue+1; j += i {
			// set primes[j] with value of i, only when primes[j] is j
			// (to prevent from overwriting previously updated lowest prime number)
			if primes[j] == j {
				primes[j] = i
			}
		}
	}

	for _, num := range nums {
		primeFactors := []int{}

		for num > 1 {
			primeFactors = append(primeFactors, primes[num])
			num /= primes[num]
		}

		ans = append(ans, primeFactors)
	}

	return ans
}

/*
		  0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18
primes = [0 1 2 3 2 5 2 7 2 3  2 11  2 13  2  3  2 17  2]
            ans
value = 12         primes[12] = 2 ans.push(2) [2]
value = 12 / 2 = 6 primes[6] = 2 ans.push(2) [2, 2]
value =  6 / 2 = 3 primes[3] = 3 ans.push(3) [2, 2, 3]
*/
