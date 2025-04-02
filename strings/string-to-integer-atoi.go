package main

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	ans := 0
	s = strings.TrimSpace(s)
	multiplier := 1

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		multiplier = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, v := range s {
		if !unicode.IsDigit(v) {
			break
		}

		current := int(v - '0')

		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && current >= 7) {
			return math.MaxInt32
		}

		if ans < math.MinInt32/10 || (ans == math.MinInt32/10 && current >= 8) {
			return math.MinInt32
		}

		ans = ans*10 + current*multiplier
	}

	return ans
}
