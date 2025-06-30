package main

import "math"

func pageFault(capacity int, pages []int) int {
	memory := map[int]int{}
	ans := 0

	for index, page := range pages {
		// put the pages within capacity and count the answer
		if index < capacity {
			memory[page] = index
			ans++
		} else { // capacity exceed
			// check if we have the item present already and replace the index
			if _, ok := memory[page]; ok {
				// update existing index here
				memory[page] = index
			} else { // item not present, remove the least item from the memory and add this new page and count
				// remove least and add new
				least := math.MaxInt
				leastPage := 0

				for p, i := range memory {
					if i < least {
						least = i
						leastPage = p
					}
				}

				delete(memory, leastPage)
				memory[page] = index

				ans++
			}
		}
	}

	return ans
}
