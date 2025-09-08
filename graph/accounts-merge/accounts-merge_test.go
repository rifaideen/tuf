package main

import "testing"

func Test_accountsMerge(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		accounts [][]string
		want     [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := accountsMerge(tt.accounts)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("accountsMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
