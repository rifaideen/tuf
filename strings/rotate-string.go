package main

import "strings"

// check if goal can be achieved after rotating the string for some number of times
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	concatenated := s + s

	return strings.Contains(concatenated, goal)
}
