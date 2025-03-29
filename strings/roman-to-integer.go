package main

func romanToInt(s string) int {
	ans := 0
	charMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"XC": 90,
		"L":  50,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	// starting from right to left
	for i := len(s) - 1; i >= 0; i-- {
		// check if we go back one step and lookup for special case
		if i-1 >= 0 {
			lookup := s[i-1 : i+1]
			// special case found, no need to go for previous value again
			if val, ok := charMap[lookup]; ok {
				ans += val
				i--

				continue
			}
		}

		// either special case not found or we can't 1 step back
		ans += charMap[string(s[i])]
	}

	return ans
}
