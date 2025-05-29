package main

// Brute-force solution
func FindCelebrity(mat [][]int) int {
	n := len(mat)
	iKnow := make([]int, n)      // Count of people 'i' knows (excluding self)
	whoKnowsMe := make([]int, n) // Count of people who know 'j' (excluding 'j' knowing self)

	for i := range n {
		for j := range n {
			// If i knows j AND i is not j (to handle the common interpretation of celebrity problem)
			if mat[i][j] == 1 && i != j {
				iKnow[i]++
				whoKnowsMe[j]++
			}
		}
	}

	for i := range n {
		// A person is a celebrity if:
		// 1. They don't know anyone else (iKnow[i] == 0)
		// 2. Everyone else knows them (whoKnowsMe[i] == n-1)
		if iKnow[i] == 0 && whoKnowsMe[i] == n-1 {
			return i
		}
	}

	return -1
}

func FindCelebrityOptimized(mat [][]int) int {
	n := len(mat)
	top := 0
	bottom := n - 1

	for top < bottom {
		if mat[top][bottom] == 1 {
			top++
		} else if mat[bottom][top] == 1 {
			bottom--
		} else {
			top++
			bottom--
		}
	}

	if top > bottom {
		return -1
	}

	for i := range n {
		if i == top {
			continue
		}

		if mat[top][i] == 0 && mat[i][top] == 1 {
			continue
		} else {
			return -1
		}
	}

	return top
}
