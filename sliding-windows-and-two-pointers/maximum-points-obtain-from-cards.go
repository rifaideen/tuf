package main

func maxScore(cardPoints []int, k int) int {
	sum := 0

	// calculate sum upto k
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	// sum is the max point we can earn so far
	maxSum := sum

	// right pointer, poiting the end
	right := len(cardPoints) - 1

	// from kth position, we subtract it from sum and add the value of right
	for i := k - 1; i >= 0; i-- {
		// remove the value from k th position in decreasing order
		sum -= cardPoints[i]
		// add the current right value to the index
		sum += cardPoints[right]

		// check if we can get new max sum
		maxSum = max(maxSum, sum)
		right--
	}

	return maxSum
}
