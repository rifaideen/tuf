package main

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	reversed := strings.Split(s, " ")
	ans := []string{}

	for i := len(reversed) - 1; i >= 0; i-- {
		if reversed[i] != "" {
			ans = append(ans, reversed[i])
		}
	}

	return strings.Join(ans, " ")
}

func reverseWordsSimplified(s string) string {
	reversed := strings.Fields(s)
	slices.Reverse(reversed)

	return strings.Join(reversed, " ")
}
