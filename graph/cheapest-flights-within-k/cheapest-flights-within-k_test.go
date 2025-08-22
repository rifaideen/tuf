package main

import "testing"

func Test_findCheapestPrice(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("findCheapestPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
