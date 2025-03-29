package main

func isAnagram(s string, t string) bool {
	strMap := map[rune]int{}

	// count the frequency of charecters in s
	for _, v := range s {
		strMap[v] += 1
	}

	// deduct the frequency of charecters in t
	for _, v := range t {
		strMap[v] -= 1
	}

	// check if we have str map is zero
	for _, v := range strMap {
		if v != 0 {
			return false
		}
	}

	return true
}
