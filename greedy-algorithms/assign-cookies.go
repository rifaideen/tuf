package main

import "slices"

func findContentChildren(g []int, s []int) int {
	// since the student at index i want to eat cookie >= student at index i
	// we can solve by sorting the students and cookies
	slices.Sort(g)
	slices.Sort(s)

	// create two pointer
	student := 0
	cookie := 0

	m := len(g)
	n := len(s)

	// while pointers are within range
	for student < m && cookie < n {
		// check if cookie is >= current value set by the student
		if s[cookie] >= g[student] {
			// requirements met, move the student
			student++
		}

		cookie++
	}

	return student
}
