package main

import "testing"

func Test_lemonadeChange(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bills []int
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lemonadeChange(tt.bills)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
