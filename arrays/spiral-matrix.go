package main

// Create 4 variable to define each corner of the matrix i.e left, right, top, bottom
//
// while top <= bottom and left <= right
//
//	traverse from left to right, increment the top.
//
//	traverse from top to bottom, decrement the right.
//
//	traverse from right to left, decrement the bottom.
//
//	traverse from bottom to top, increment the left.
func SpiralMatrix(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	answer := []int{}

	// create 4 variables to define each corner of the matrix with initial values
	top := 0
	bottom := rows - 1
	left := 0
	right := cols - 1

	// as we go through the matrix we update each corners
	// while top <= bottom and left <= right, extract the information from matrix
	for top <= bottom && left <= right {
		// traverse from left to right and then move the top by 1 after the loop
		for i := left; i <= right; i++ {
			answer = append(answer, matrix[top][i])
		}

		top++

		// traverse from top to bottom and move right by -1 after the loop
		for i := top; i <= bottom; i++ {
			answer = append(answer, matrix[i][right])
		}

		right--

		// traverse from right to left and move bottom by -1, after the loop
		if top <= bottom { // since we are using bottom index, make sure it is within the range
			for i := right; i >= left; i-- {
				answer = append(answer, matrix[bottom][i])
			}

			bottom--
		}

		// traverse from bottom to top and move left by 1
		if left <= right { // since we are using left index, make sure it is withing the range
			for i := bottom; i >= top; i-- {
				answer = append(answer, matrix[i][left])
			}

			left++
		}
	}

	return answer
}
