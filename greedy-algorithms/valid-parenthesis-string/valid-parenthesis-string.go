package main

import "fmt"

// check valid parenthesis string solves the problem using recursive approach
func checkValidString(s string) bool {
	if len(s) == 1 {
		return s == "*"
	}

	// create a map which contains precomputation result to avoid expensive calculation again
	memo := map[string]bool{}

	return checkValidStringRecursive(0, len(s), 0, s, &memo)
}

func checkValidStringRecursive(index int, n int, sum int, sequence string, cache *map[string]bool) bool {
	// create unique key
	key := fmt.Sprintf("%d.%d", index, sum)

	// at any time, sum becomes, we update the case and return false
	if sum < 0 {
		(*cache)[key] = false

		return false
	}

	// if we already computed, return the computed result from cache
	if val, ok := (*cache)[key]; ok {
		return val
	}

	// index reaches n, set the cache and return true when sum == 0
	if index == n {
		(*cache)[key] = sum == 0

		return sum == 0
	}

	switch sequence[index] {
	case '(': // increment the sum and move forward
		result := checkValidStringRecursive(index+1, n, sum+1, sequence, cache)
		(*cache)[key] = result

		return result
	case ')': // decrement the sum and move forward
		result := checkValidStringRecursive(index+1, n, sum-1, sequence, cache)
		(*cache)[key] = result

		return result
	default:
		// we have *, we can try with empty (doesn't do anything with sum)
		// we can try with open bracket, which increases sum by 1
		// we can try with close bracket, which decreases sum by 1
		empty := checkValidStringRecursive(index+1, n, sum, sequence, cache)
		open := checkValidStringRecursive(index+1, n, sum+1, sequence, cache)
		close := checkValidStringRecursive(index+1, n, sum-1, sequence, cache)

		// store the overall result and cache it
		result := empty || open || close
		(*cache)[key] = result

		return result
	}
}

// This solves the problem using min and max range
func CheckValidStringII(s string) bool {
	if len(s) == 1 && s[0] != '*' {
		return false
	}

	minRange := 0
	maxRange := 0

	for _, i := range s {
		switch i {
		case '(':
			minRange++
			maxRange++
		case ')':
			minRange--
			maxRange--
		default:
			minRange--
			maxRange++
		}

		if minRange < 0 {
			minRange = 0
		}

		if maxRange < 0 {
			return false
		}
	}

	return minRange == 0
}
