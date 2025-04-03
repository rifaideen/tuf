package main

func countSubstr(str string, k int) int {
	return atmostK(str, k) - atmostK(str, k-1)
}

func atmostK(str string, k int) int {
	ans := 0
	left := 0
	n := len(str)
	charMap := map[byte]int{}

	for right := range n {
		charMap[str[right]-'a']++

		for len(charMap) > k {
			charMap[str[left]-'a']--

			if charMap[str[left]-'a'] == 0 {
				delete(charMap, str[left]-'a')
			}

			left++
		}

		ans += right - left + 1
	}

	return ans
}
