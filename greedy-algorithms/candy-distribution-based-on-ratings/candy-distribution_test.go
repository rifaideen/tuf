package main

import "testing"

func Test_candy(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ratings []int
		want    int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := candy(tt.ratings)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
