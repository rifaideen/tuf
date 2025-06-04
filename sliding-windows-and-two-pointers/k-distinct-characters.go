package main

func kDisinctCharacters(str string, k int) int {
	left := 0
	frequency := map[byte]int{}
	ans := 0

	for right := range len(str) {
		frequency[str[right]]++

		for len(frequency) > k {
			frequency[str[left]]--

			if frequency[str[left]] == 0 {
				delete(frequency, str[left])
			}

			left++
		}

		if len(frequency) == k {
			ans = max(ans, right-left+1)
		}
	}

	return ans
}
