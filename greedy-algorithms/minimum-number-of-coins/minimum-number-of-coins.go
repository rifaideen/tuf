package main

// This particular problem solves the problem as stated in
// https://www.geeksforgeeks.org/problems/-minimum-number-of-coins4426/1
func minPartition(n int) []int {
	rupees := []int{2000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

	change := []int{}

	for _, rupee := range rupees {
		if n >= rupee {
			for range n / rupee {
				change = append(change, rupee)
				n -= rupee
			}
		}
	}

	return change
}
