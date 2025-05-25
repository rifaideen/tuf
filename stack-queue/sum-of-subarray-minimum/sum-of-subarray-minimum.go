package main

func sumSubarrayMins(arr []int) int {
	mod := int(1000_000_000) + 7
	pse := pse(arr)
	nse := nse(arr)
	total := 0

	for i, val := range arr {
		left := i - pse[i]
		right := nse[i] - i

		// multiplying left and right give us number of contributions
		// multiplying that with current value gives us sum of subarray minimum for the current element
		total = total + ((left * right * val) % mod)
		total %= mod
	}

	return total
}

// find next smaller element for each item in array
func nse(arr []int) []int {
	n := len(arr)
	nse := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}

		nse[i] = n

		if len(stack) > 0 {
			nse[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return nse
}

// find previous smaller element for each item in array
func pse(arr []int) []int {
	n := len(arr)
	pse := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}

		pse[i] = -1

		if len(stack) > 0 {
			pse[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return pse
}
