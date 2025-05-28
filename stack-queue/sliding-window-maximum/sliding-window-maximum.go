package main

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	n := len(nums)

	// consider it dqueue where I can push/pop on front and back
	// stores in decreasing order (decreasing from front i.e 5,4,3,2,1)
	// It store the indices whose values are in decreasing order
	dqueue := []int{}

	for i := 0; i < n; i++ {
		// check if front exceed the window and shrink it
		if len(dqueue) > 0 && dqueue[0] < i-k+1 {
			dqueue = dqueue[1:]
		}

		// since we need to store the values in decreasing order
		// while dqueue's top is less than current value, pop it out continously
		for len(dqueue) > 0 && nums[dqueue[len(dqueue)-1]] <= nums[i] {
			dqueue = dqueue[:len(dqueue)-1]
		}

		// push the current index to dqueue
		dqueue = append(dqueue, i)

		// push it to anwer only when we started filling out the window size
		// i.e for k = 3, the moment i crosses k-1, we can start to populate the answer
		// from the dqueue front
		if i >= k-1 {
			ans = append(ans, nums[dqueue[0]])
		}
	}

	return ans
}
