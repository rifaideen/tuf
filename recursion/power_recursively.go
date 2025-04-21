package main

func myPow(x float64, n int) float64 {
	// anything power zero is 1
	if n == 0 {
		return 1
	}

	// zero power anything is 0
	if x == 0 {
		return 0
	}

	// check if n is negative
	if n < 0 {
		x = 1 / x // apply the formula pow = 1/x^n
		n = -n    // convert the n into positive
	}

	if n%2 == 0 { // handle even number of n with x2^(n/2)
		return myPow(x*x, n/2)
	}

	// handle odd case x * (x2 ^ ((n-1)/2))
	return x * myPow(x*x, (n-1)/2)
}
