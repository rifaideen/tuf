package main

// letterCombinations generates all possible letter combinations from the given digits
// Each digit maps to multiple letters (like on a phone keypad)
func letterCombinations(digits string) []string {
	ans := []string{} // Initialize result array to store all valid combinations
	n := len(digits)  // Store the length of input digits

	// Call the recursive helper function to generate combinations
	letterCombinationsRecursive(0, n, &ans, &[]rune{}, digits)
	return ans
}

// letterCombinationsRecursive is a backtracking function that builds combinations recursively
// index: current position in the digits string we're processing
// n: total length of digits string
// ans: pointer to the result array where valid combinations are stored
// sequence: pointer to the current combination being built
// digits: the input string of digits
func letterCombinationsRecursive(index int, n int, ans *[]string, sequence *[]rune, digits string) {
	// Base case: if we've processed all digits (reached the end of input)
	if index == n {
		// Only add non-empty sequences to the result
		if len(*sequence) > 0 {
			*ans = append(*ans, string(*sequence))
		}
		return
	}

	// Define the mapping from digits to letters (phone keypad)
	mapping := map[int][]rune{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}

	// Get all possible letters for the current digit
	chars := mapping[int(digits[index]-'0')]

	// Loop-based backtracking approach:
	// Try each possible letter for the current position
	for _, char := range chars {
		// 1. ADD: Add the current character to our sequence
		*sequence = append(*sequence, char)

		// 2. EXPLORE: Recursively process the next digit with updated sequence
		letterCombinationsRecursive(index+1, n, ans, sequence, digits)

		// 3. REMOVE (BACKTRACK): Remove the last character to try the next possibility
		// This is crucial for backtracking - undo the choice before trying the next option
		*sequence = (*sequence)[:len(*sequence)-1]

		// Note: We don't make a second recursive call here because:
		// - This is a loop-based approach, not a "pick/don't pick" approach
		// - Each position must have exactly one letter from the current digit
		// - The loop handles trying all possibilities systematically
	}
}
