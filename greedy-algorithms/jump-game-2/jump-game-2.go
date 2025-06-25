package main

// we solve this problem using the two pointers, both of them points to 0th position
// and left pointer moves 1 step after right and right pointers points the highest index
// we can reach within the left and right pointer.
// we keep track how variable which counts the number of jumps inside the loop and returned
func jump(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	jumps := 0

	// create two pointers, both pointing to 0th index initially
	left := 0
	right := 0

	// the moment, right pointer reaches the last item, we stop the loop
	for right < n-1 {
		maxrange := 0

		// find the maximum index we can reach between the pointers
		// this helps us update the pointers position after this
		//
		// why we did not choose the maximum from left or right is, within range,
		// there could be a bigger step which could leads to smaller steps.
		//
		// that is the reason we start the loop from left to right to try out all possibilities
		for i := left; i <= right; i++ {
			maxrange = max(maxrange, i+nums[i])
		}

		// left pointer always move 1 step after right
		left = right + 1
		// right pointers always point out the maximum index we can travel
		right = maxrange
		// we have completed 1 jump within the pointers
		jumps++
	}

	return jumps
}
