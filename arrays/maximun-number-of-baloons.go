package main

import "math"

// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
//
// You can use each character in text at most once. Return the maxilmum number of instances that can be formed.
//
// Intuition:
// 1, we calculare the pattern frequency i.e balloon is the pattern with frequency of
// patternFrequency = {b:1, a:1, l:2, o:2, n:1}
// 2, calculate the frequency for the given string str i.e loonbaxballpoon
// strFrequency = {l:4, o:4, n:2, b:2, a:2, x:1, p:1 }
// 3, calculate answer by iterating from 0 - 26 (0-25 inclusive) and check if the value present in patternFrequency
// answer = min(answer, strFrequency[i] / patternFrequency[i])
// 4, return the answer
func MaxBaloons(str string) int {
	pattern := "balloon"

	strFrequency := make([]int, 26) // set capacity of 26 to store a - z
	patternFrequency := make([]int, 26)

	for _, s := range str {
		strFrequency[s-'a']++
	}

	for _, p := range pattern {
		patternFrequency[p-'a']++
	}

	answer := math.MaxInt

	for i := 0; i < 26; i++ {
		if patternFrequency[i] > 0 {
			f := strFrequency[i] / patternFrequency[i]
			answer = min(answer, f)
		}
	}

	return answer
}
