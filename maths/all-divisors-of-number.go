package main

func divisors(n int) []int {
	ans := []int{}

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ans = append(ans, i)

			if n/i != i {
				ans = append(ans, n/i)
			}
		}
	}

	return ans
}
