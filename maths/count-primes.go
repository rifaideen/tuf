package main

func countPrimes(n int) int {
	ans := make([]int, n)

	// fill the ans array with 1 (consider everything being prime number at the moment)
	for i := range n {
		ans[i] = 1
	}

	// start from the least prime number to the n and see if we have prime number set in the current index
	for i := 2; i < n; i++ {
		if ans[i] == 1 {
			// mark all the multiple of i to 0
			// we are resetting the 2's multiples and 3's multiples to 0 and any number which are multiple of i are being set to 0
			// so we ultimately left with whatever is prime being untouched
			for j := i * i; j < n; j += i {
				ans[j] = 0
			}
		}
	}

	sum := 0

	// at this point, we marked the prime numbers and we just need to count them from the least prime number
	for i := 2; i < n; i++ {
		sum += ans[i]
	}

	return sum
}
