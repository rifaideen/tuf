package main

func printAllPrimeFactors(n int) []int {
	ans := []int{}

	// we loop from 2 to square root of n
	for i := 2; i*i <= n; i++ {
		// if i can be perfectly divide the n, we reduce the n until this condition breaks
		if n%i == 0 {
			// push it to ans
			ans = append(ans, i)

			for n%i == 0 {
				n /= i
			}
		}
	}

	// in case n is not reduces at all, we need to include the current value of n
	if n != 1 {
		ans = append(ans, n)
	}

	return ans
}
