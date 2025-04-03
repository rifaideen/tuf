package main

func longestPalindrome(s string) string {
	res := ""
	ans := 0

	for i := range len(s) {
		// for each iteration of i, peform palindrome check for odd and even case
		left := i
		right := i

		// odd case
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > ans {
				res = s[left : right+1]
				ans = right - left + 1
			}

			left--
			right++
		}

		left = i
		right = i + 1

		// even case
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > ans {
				res = s[left : right+1]
				ans = right - left + 1
			}

			left--
			right++
		}
	}

	return res
}
