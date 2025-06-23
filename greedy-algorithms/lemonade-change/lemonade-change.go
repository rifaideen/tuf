package main

func lemonadeChange(bills []int) bool {
	// creating array of size 3 to hold 5,10,20 bills
	changes := [3]int{
		0, // 5$
		0, // 10$
		0, // 20$
	}

	for _, bill := range bills {
		if bill == 5 {
			changes[0]++
		} else if bill == 10 {
			if changes[0] == 0 { // I have no changes, stop here
				return false
			}

			changes[0]--
			changes[1]++
		} else {
			// I have enough 10 and 5
			if changes[1] >= 1 && changes[0] >= 1 {
				changes[0]-- // give 10 + 5
				changes[1]--
				changes[2]++
			} else if changes[0] >= 3 { // I have no 10s but enough 5s
				changes[0] -= 3
				changes[2]++
			} else { // I don't have enough changes
				return false
			}
		}
	}

	return true
}
