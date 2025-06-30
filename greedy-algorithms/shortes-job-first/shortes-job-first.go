package main

import "slices"

func averageTimes(jobs []int) int {
	n := len(jobs)

	// sort the jobs and start processing jobs with less time to complete the task
	slices.Sort(jobs)

	// current time when the jobs arrived
	t := 0
	// accumulated waiting time to complete the jobs
	waiting_time := 0

	for _, job := range jobs {
		waiting_time += t
		t += job
	}

	// find the average waiting time
	return waiting_time / n
}
