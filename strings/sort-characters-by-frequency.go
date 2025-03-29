package main

import (
	"cmp"
	"slices"
	"strings"
)

func frequencySort(s string) string {
	frequency := map[rune]int{}

	// count the frequency of charecters in s
	for _, char := range s {
		frequency[char]++
	}

	// create set to store the char and it's frequency
	set := []any{}

	for char, val := range frequency {
		set = append(set, []any{string(char), val})
	}

	// sort the set by frequency
	slices.SortFunc(set, func(a, b any) int {
		return cmp.Compare(b.([]any)[1].(int), a.([]any)[1].(int))
	})

	ans := ""

	// loop the set and create the same number of charecter as specified in each set
	for _, char := range set {
		ans += strings.Repeat(char.([]any)[0].(string), char.([]any)[1].(int))
	}

	return ans
}
