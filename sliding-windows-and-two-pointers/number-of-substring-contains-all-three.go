package main

func numberOfSubstrings(s string) int {
	// to keep track of where we saw the a,b and c (the last seen index)
	frequency := [3]int{-1, -1, -1}
	right := 0
	ans := 0

	for right < len(s) {
		// update the last seen index
		frequency[s[right]-'a'] = right

		// if all of our last seen index are not empty ( > -1)
		// we can easily count how many number of valid sub-arrays ending at current iteration
		// we then identify the minimum last seen index among the frequency array, increment it by 1
		// and add it to existing answer
		// this way, the TC is still O(N)
		if frequency[0] > -1 && frequency[1] > -1 && frequency[2] > -1 {
			ans += min(frequency[0], frequency[1], frequency[2]) + 1
		}

		right++
	}

	return ans
}
