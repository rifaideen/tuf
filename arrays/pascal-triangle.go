package main

// find nCr for n row and r col in pascal triangle using formula nCr = n! / (r! * (n-r)!)
func ncr(n, r int) int {
	ans := 1

	for i := range r {
		ans *= (n - i)
		ans /= (i + 1)
	}

	return ans
}

// GetPascalElement returns the element at row, col in pascal triangle which is variation #1
func GetPascalElement(row, col int) int {
	return ncr(row-1, col-1)
}

// GetPascalRow returns the row at row in pascal triangle which is variation #2
// 0 based index row
func GetPascalRow(row int) []int {
	ans := 1          // start with 1
	res := []int{ans} // start the result with initial answer

	// find ncr and append the result `res`
	for i := range row {
		ans *= (row - i)
		ans /= (i + 1)

		res = append(res, ans)
	}

	return res
}

func GetPascalTriangle(rows int) [][]int {
	res := [][]int{}

	for i := range rows {
		res = append(res, GetPascalRow(i))
	}

	return res
}
