package main

// we solve this problem using the efficient algorithm by thinking of provided ratings as slope
// when the slope increases, we increment the peak by 1 and assign
// when the slop decreases, we reset the candy to 1 and increment by 1 and assingn
// once done, we compare the peak and down value, in case down is greater than peak, we add the difference to the sum
func candy(ratings []int) int {
	sum := 1
	// candy := 1
	i := 1
	n := len(ratings)

	for i < n {
		// handles flat slope and assign 1 candy to the students
		if ratings[i] == ratings[i-1] {
			sum += 1
			i++
			continue
		}

		// handles increasing slope
		peak := 1

		for i < n && ratings[i] > ratings[i-1] {
			peak += 1
			sum += peak
			i++
		}

		down := 1

		// handles decreasing slope
		for i < n && ratings[i] < ratings[i-1] {
			sum += down
			down += 1
			i++
		}

		if down > peak {
			sum += down - peak
		}
	}

	return sum
}
