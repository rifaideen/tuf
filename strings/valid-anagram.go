package main

func isAnagram(s string, t string) bool {
	strMap := map[rune]int{}

	for _, v := range s {
		strMap[v] += 1
	}

	for _, v := range t {
		strMap[v] -= 1
	}

	for _, v := range strMap {
		if v != 0 {
			return false
		}
	}

	return true
}
