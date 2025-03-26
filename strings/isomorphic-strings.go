package main

func isIsomorphic(s string, t string) bool {
	// we keep two map to store the actual char to replacement char
	map1 := map[byte]byte{}
	map2 := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		value1 := s[i]
		value2 := t[i]

		// check if map1 for value1 contains value2 as the replacement
		if val, ok := map1[value1]; ok && val != value2 {
			return false
		} else {
			// add value2 as the replacement for value1
			map1[value1] = value2
		}

		if val, ok := map2[value2]; ok && val != value1 {
			return false
		} else {
			// add value1 as the replacement for value2
			map2[value2] = value1
		}
	}

	// the for loop terminated, we found matching charecters, hence isomorphic
	return true
}
