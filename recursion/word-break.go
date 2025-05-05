package main

func wordBreak(s string, wordDict []string) bool {
	dictionary := map[string]bool{}

	for _, word := range wordDict {
		dictionary[word] = true
	}

	memo := map[int]bool{}

	n := len(s)

	return wordBreakRecursive(0, n, dictionary, s, memo)
}

func wordBreakRecursive(index, n int, dictionary map[string]bool, s string, memo map[int]bool) bool {
	if index == n {
		return true
	}

	if val, ok := memo[index]; ok {
		return val
	}

	for i := index; i < n; i++ {
		if _, ok := dictionary[string(s[index:i+1])]; ok {
			if wordBreakRecursive(i+1, n, dictionary, s, memo) {
				memo[index] = true
				return true
			}
		}
	}

	memo[index] = false
	return false
}
