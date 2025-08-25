package main

import "testing"

func Test_networkDelayTime(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		times [][]int
		n     int
		k     int
		want  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := networkDelayTime(tt.times, tt.n, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
