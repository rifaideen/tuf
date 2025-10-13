package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unboundedKnapsack(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		val     []int
		n       int
		W       int
		want    int
	}{
		{
			name:    "single_item_fit_perfectly",
			weights: []int{10},
			val:     []int{20},
			n:       1,
			W:       10,
			want:    20,
		},
		{
			name:    "single_item_does_not_fit",
			weights: []int{15},
			val:     []int{30},
			n:       1,
			W:       10,
			want:    0,
		},
		{
			name:    "single_item_multiple_copies_possible",
			weights: []int{5},
			val:     []int{10},
			n:       1,
			W:       15,
			want:    30, // 3 copies: 5*3 <= 15 -> value = 10*3 = 30
		},
		{
			name:    "two_items_equal_weight_higher_value_better",
			weights: []int{5, 5},
			val:     []int{10, 15},
			n:       2,
			W:       10,
			want:    30, // take item 2 twice: 15 * 2 = 30
		},
		{
			name:    "classic_small_case_optimal_combination",
			weights: []int{2, 3, 5},
			val:     []int{10, 15, 30},
			n:       3,
			W:       8,
			want:    45, // one of weight 2 (value 10) + two of weight 3 (value 15*2=30) => total 45
			// or three of weight 2 and one of weight 2 again? Let's verify:
			// Actually best is: one wt=2 (10), two wt=3 (30) → total wt=8, val=45
			// Alternatively: one wt=5 (30) + one wt=3 (15) → total 8, val=45 → same.
		},
		{
			name:    "all_items_heavy_cannot_fit_any",
			weights: []int{600, 700},
			val:     []int{100, 200},
			n:       2,
			W:       500,
			want:    0,
		},
		{
			name:    "zero_capacity",
			weights: []int{10, 20},
			val:     []int{50, 100},
			n:       2,
			W:       0,
			want:    0,
		},
		{
			name:    "large_knapsack_with_repetition",
			weights: []int{1, 2},
			val:     []int{1, 3},
			n:       2,
			W:       1000,
			want:    1500, // always pick item with higher value per weight: val/wt = 3/2 = 1.5 > 1
			// so use as many of weight 2 as possible: 1000 / 2 = 500 items → 500 * 3 = 1500
		},
		{
			name:    "better_density_but_not_divisible_weight",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Corrected from 26 → 25
		},
		// Fixing above with correct expected value
		{
			name:    "corrected_two_item_greedy_not_optimal",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Two of weight 3 (6, value 16) + one of weight 4 (4, value 9) → total wt=10, val=25
			// Or: one of weight 3 (3, 8) + two of weight 4 (8, 18) → total 11 → invalid
			// So maximum is 25.
		},
		{
			name:    "repeat_best_item_multiple_times_optimal",
			weights: []int{5, 10},
			val:     []int{10, 19},
			n:       2,
			W:       20,
			want:    40, // Two items of weight 5 (each value 10): 20/5 = 4 → 4*10 = 40
			// vs one item of weight 10: 20/10 = 2 → 2*19 = 38 < 40
		},
		{
			name:    "high_value_low_weight_dominates",
			weights: []int{1, 500},
			val:     []int{500, 1},
			n:       2,
			W:       1000,
			want:    500000, // take 1000 copies of first item: 1000 * 500 = 500,000
		},
		{
			name:    "multiple_same_weight_different_values",
			weights: []int{10, 10, 10},
			val:     []int{10, 20, 30},
			n:       3,
			W:       25,
			want:    60, // Corrected from 75 → 60
		},
		{
			name:    "exact_fit_with_max_repetition",
			weights: []int{10, 20},
			val:     []int{30, 50},
			n:       2,
			W:       30,
			want:    90, // three of first item: 3 * 30 = 90
			// vs one of each: 30+50=80 → less than 90
		},
		{
			name:    "worst_value_per_weight_but_fits_perfectly",
			weights: []int{500},
			val:     []int{1},
			n:       1,
			W:       1000,
			want:    2, // two copies possible
		},
		{
			name:    "many_items_with_small_weights_and_values",
			weights: []int{1, 2, 3},
			val:     []int{1, 2, 3},
			n:       3,
			W:       10,
			want:    10, // all have same value/weight ratio (1), so any combo summing to 10 gives 10
		},
		{
			name:    "best_ratio_in_middle_of_array",
			weights: []int{10, 1, 5},
			val:     []int{10, 2, 7},
			// ratios: 1.0, 2.0, 1.4 → best is index 1 (val=2, wt=1)
			n:    3,
			W:    7,
			want: 14, // 7 copies of item 1 → 7 * 2 = 14
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unboundedKnapsack(tt.weights, tt.val, tt.n, tt.W)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_unboundedKnapsackMemoized(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		val     []int
		n       int
		W       int
		want    int
	}{
		{
			name:    "single_item_fit_perfectly",
			weights: []int{10},
			val:     []int{20},
			n:       1,
			W:       10,
			want:    20,
		},
		{
			name:    "single_item_does_not_fit",
			weights: []int{15},
			val:     []int{30},
			n:       1,
			W:       10,
			want:    0,
		},
		{
			name:    "single_item_multiple_copies_possible",
			weights: []int{5},
			val:     []int{10},
			n:       1,
			W:       15,
			want:    30, // 3 copies: 5*3 <= 15 -> value = 10*3 = 30
		},
		{
			name:    "two_items_equal_weight_higher_value_better",
			weights: []int{5, 5},
			val:     []int{10, 15},
			n:       2,
			W:       10,
			want:    30, // take item 2 twice: 15 * 2 = 30
		},
		{
			name:    "classic_small_case_optimal_combination",
			weights: []int{2, 3, 5},
			val:     []int{10, 15, 30},
			n:       3,
			W:       8,
			want:    45, // one of weight 2 (value 10) + two of weight 3 (value 15*2=30) => total 45
			// or three of weight 2 and one of weight 2 again? Let's verify:
			// Actually best is: one wt=2 (10), two wt=3 (30) → total wt=8, val=45
			// Alternatively: one wt=5 (30) + one wt=3 (15) → total 8, val=45 → same.
		},
		{
			name:    "all_items_heavy_cannot_fit_any",
			weights: []int{600, 700},
			val:     []int{100, 200},
			n:       2,
			W:       500,
			want:    0,
		},
		{
			name:    "zero_capacity",
			weights: []int{10, 20},
			val:     []int{50, 100},
			n:       2,
			W:       0,
			want:    0,
		},
		{
			name:    "large_knapsack_with_repetition",
			weights: []int{1, 2},
			val:     []int{1, 3},
			n:       2,
			W:       1000,
			want:    1500, // always pick item with higher value per weight: val/wt = 3/2 = 1.5 > 1
			// so use as many of weight 2 as possible: 1000 / 2 = 500 items → 500 * 3 = 1500
		},
		{
			name:    "better_density_but_not_divisible_weight",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Corrected from 26 → 25
		},
		// Fixing above with correct expected value
		{
			name:    "corrected_two_item_greedy_not_optimal",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Two of weight 3 (6, value 16) + one of weight 4 (4, value 9) → total wt=10, val=25
			// Or: one of weight 3 (3, 8) + two of weight 4 (8, 18) → total 11 → invalid
			// So maximum is 25.
		},
		{
			name:    "repeat_best_item_multiple_times_optimal",
			weights: []int{5, 10},
			val:     []int{10, 19},
			n:       2,
			W:       20,
			want:    40, // Two items of weight 5 (each value 10): 20/5 = 4 → 4*10 = 40
			// vs one item of weight 10: 20/10 = 2 → 2*19 = 38 < 40
		},
		{
			name:    "high_value_low_weight_dominates",
			weights: []int{1, 500},
			val:     []int{500, 1},
			n:       2,
			W:       1000,
			want:    500000, // take 1000 copies of first item: 1000 * 500 = 500,000
		},
		{
			name:    "multiple_same_weight_different_values",
			weights: []int{10, 10, 10},
			val:     []int{10, 20, 30},
			n:       3,
			W:       25,
			want:    60, // Corrected from 75 → 60
		},
		{
			name:    "exact_fit_with_max_repetition",
			weights: []int{10, 20},
			val:     []int{30, 50},
			n:       2,
			W:       30,
			want:    90, // three of first item: 3 * 30 = 90
			// vs one of each: 30+50=80 → less than 90
		},
		{
			name:    "worst_value_per_weight_but_fits_perfectly",
			weights: []int{500},
			val:     []int{1},
			n:       1,
			W:       1000,
			want:    2, // two copies possible
		},
		{
			name:    "many_items_with_small_weights_and_values",
			weights: []int{1, 2, 3},
			val:     []int{1, 2, 3},
			n:       3,
			W:       10,
			want:    10, // all have same value/weight ratio (1), so any combo summing to 10 gives 10
		},
		{
			name:    "best_ratio_in_middle_of_array",
			weights: []int{10, 1, 5},
			val:     []int{10, 2, 7},
			// ratios: 1.0, 2.0, 1.4 → best is index 1 (val=2, wt=1)
			n:    3,
			W:    7,
			want: 14, // 7 copies of item 1 → 7 * 2 = 14
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unboundedKnapsackMemoized(tt.weights, tt.val, tt.n, tt.W)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_unboundedKnapsackTabulation(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		val     []int
		n       int
		W       int
		want    int
	}{
		{
			name:    "single_item_fit_perfectly",
			weights: []int{10},
			val:     []int{20},
			n:       1,
			W:       10,
			want:    20,
		},
		{
			name:    "single_item_does_not_fit",
			weights: []int{15},
			val:     []int{30},
			n:       1,
			W:       10,
			want:    0,
		},
		{
			name:    "single_item_multiple_copies_possible",
			weights: []int{5},
			val:     []int{10},
			n:       1,
			W:       15,
			want:    30, // 3 copies: 5*3 <= 15 -> value = 10*3 = 30
		},
		{
			name:    "two_items_equal_weight_higher_value_better",
			weights: []int{5, 5},
			val:     []int{10, 15},
			n:       2,
			W:       10,
			want:    30, // take item 2 twice: 15 * 2 = 30
		},
		{
			name:    "classic_small_case_optimal_combination",
			weights: []int{2, 3, 5},
			val:     []int{10, 15, 30},
			n:       3,
			W:       8,
			want:    45, // one of weight 2 (value 10) + two of weight 3 (value 15*2=30) => total 45
			// or three of weight 2 and one of weight 2 again? Let's verify:
			// Actually best is: one wt=2 (10), two wt=3 (30) → total wt=8, val=45
			// Alternatively: one wt=5 (30) + one wt=3 (15) → total 8, val=45 → same.
		},
		{
			name:    "all_items_heavy_cannot_fit_any",
			weights: []int{600, 700},
			val:     []int{100, 200},
			n:       2,
			W:       500,
			want:    0,
		},
		{
			name:    "zero_capacity",
			weights: []int{10, 20},
			val:     []int{50, 100},
			n:       2,
			W:       0,
			want:    0,
		},
		{
			name:    "large_knapsack_with_repetition",
			weights: []int{1, 2},
			val:     []int{1, 3},
			n:       2,
			W:       1000,
			want:    1500, // always pick item with higher value per weight: val/wt = 3/2 = 1.5 > 1
			// so use as many of weight 2 as possible: 1000 / 2 = 500 items → 500 * 3 = 1500
		},
		{
			name:    "better_density_but_not_divisible_weight",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Corrected from 26 → 25
		},
		// Fixing above with correct expected value
		{
			name:    "corrected_two_item_greedy_not_optimal",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Two of weight 3 (6, value 16) + one of weight 4 (4, value 9) → total wt=10, val=25
			// Or: one of weight 3 (3, 8) + two of weight 4 (8, 18) → total 11 → invalid
			// So maximum is 25.
		},
		{
			name:    "repeat_best_item_multiple_times_optimal",
			weights: []int{5, 10},
			val:     []int{10, 19},
			n:       2,
			W:       20,
			want:    40, // Two items of weight 5 (each value 10): 20/5 = 4 → 4*10 = 40
			// vs one item of weight 10: 20/10 = 2 → 2*19 = 38 < 40
		},
		{
			name:    "high_value_low_weight_dominates",
			weights: []int{1, 500},
			val:     []int{500, 1},
			n:       2,
			W:       1000,
			want:    500000, // take 1000 copies of first item: 1000 * 500 = 500,000
		},
		{
			name:    "multiple_same_weight_different_values",
			weights: []int{10, 10, 10},
			val:     []int{10, 20, 30},
			n:       3,
			W:       25,
			want:    60, // Corrected from 75 → 60
		},
		{
			name:    "exact_fit_with_max_repetition",
			weights: []int{10, 20},
			val:     []int{30, 50},
			n:       2,
			W:       30,
			want:    90, // three of first item: 3 * 30 = 90
			// vs one of each: 30+50=80 → less than 90
		},
		{
			name:    "worst_value_per_weight_but_fits_perfectly",
			weights: []int{500},
			val:     []int{1},
			n:       1,
			W:       1000,
			want:    2, // two copies possible
		},
		{
			name:    "many_items_with_small_weights_and_values",
			weights: []int{1, 2, 3},
			val:     []int{1, 2, 3},
			n:       3,
			W:       10,
			want:    10, // all have same value/weight ratio (1), so any combo summing to 10 gives 10
		},
		{
			name:    "best_ratio_in_middle_of_array",
			weights: []int{10, 1, 5},
			val:     []int{10, 2, 7},
			// ratios: 1.0, 2.0, 1.4 → best is index 1 (val=2, wt=1)
			n:    3,
			W:    7,
			want: 14, // 7 copies of item 1 → 7 * 2 = 14
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unboundedKnapsackTabulation(tt.weights, tt.val, tt.n, tt.W)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_unboundedKnapsackOptimized(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		val     []int
		n       int
		W       int
		want    int
	}{
		{
			name:    "single_item_fit_perfectly",
			weights: []int{10},
			val:     []int{20},
			n:       1,
			W:       10,
			want:    20,
		},
		{
			name:    "single_item_does_not_fit",
			weights: []int{15},
			val:     []int{30},
			n:       1,
			W:       10,
			want:    0,
		},
		{
			name:    "single_item_multiple_copies_possible",
			weights: []int{5},
			val:     []int{10},
			n:       1,
			W:       15,
			want:    30, // 3 copies: 5*3 <= 15 -> value = 10*3 = 30
		},
		{
			name:    "two_items_equal_weight_higher_value_better",
			weights: []int{5, 5},
			val:     []int{10, 15},
			n:       2,
			W:       10,
			want:    30, // take item 2 twice: 15 * 2 = 30
		},
		{
			name:    "classic_small_case_optimal_combination",
			weights: []int{2, 3, 5},
			val:     []int{10, 15, 30},
			n:       3,
			W:       8,
			want:    45, // one of weight 2 (value 10) + two of weight 3 (value 15*2=30) => total 45
			// or three of weight 2 and one of weight 2 again? Let's verify:
			// Actually best is: one wt=2 (10), two wt=3 (30) → total wt=8, val=45
			// Alternatively: one wt=5 (30) + one wt=3 (15) → total 8, val=45 → same.
		},
		{
			name:    "all_items_heavy_cannot_fit_any",
			weights: []int{600, 700},
			val:     []int{100, 200},
			n:       2,
			W:       500,
			want:    0,
		},
		{
			name:    "zero_capacity",
			weights: []int{10, 20},
			val:     []int{50, 100},
			n:       2,
			W:       0,
			want:    0,
		},
		{
			name:    "large_knapsack_with_repetition",
			weights: []int{1, 2},
			val:     []int{1, 3},
			n:       2,
			W:       1000,
			want:    1500, // always pick item with higher value per weight: val/wt = 3/2 = 1.5 > 1
			// so use as many of weight 2 as possible: 1000 / 2 = 500 items → 500 * 3 = 1500
		},
		{
			name:    "better_density_but_not_divisible_weight",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Corrected from 26 → 25
		},
		// Fixing above with correct expected value
		{
			name:    "corrected_two_item_greedy_not_optimal",
			weights: []int{3, 4},
			val:     []int{8, 9},
			n:       2,
			W:       10,
			want:    25, // Two of weight 3 (6, value 16) + one of weight 4 (4, value 9) → total wt=10, val=25
			// Or: one of weight 3 (3, 8) + two of weight 4 (8, 18) → total 11 → invalid
			// So maximum is 25.
		},
		{
			name:    "repeat_best_item_multiple_times_optimal",
			weights: []int{5, 10},
			val:     []int{10, 19},
			n:       2,
			W:       20,
			want:    40, // Two items of weight 5 (each value 10): 20/5 = 4 → 4*10 = 40
			// vs one item of weight 10: 20/10 = 2 → 2*19 = 38 < 40
		},
		{
			name:    "high_value_low_weight_dominates",
			weights: []int{1, 500},
			val:     []int{500, 1},
			n:       2,
			W:       1000,
			want:    500000, // take 1000 copies of first item: 1000 * 500 = 500,000
		},
		{
			name:    "multiple_same_weight_different_values",
			weights: []int{10, 10, 10},
			val:     []int{10, 20, 30},
			n:       3,
			W:       25,
			want:    60, // Corrected from 75 → 60
		},
		{
			name:    "exact_fit_with_max_repetition",
			weights: []int{10, 20},
			val:     []int{30, 50},
			n:       2,
			W:       30,
			want:    90, // three of first item: 3 * 30 = 90
			// vs one of each: 30+50=80 → less than 90
		},
		{
			name:    "worst_value_per_weight_but_fits_perfectly",
			weights: []int{500},
			val:     []int{1},
			n:       1,
			W:       1000,
			want:    2, // two copies possible
		},
		{
			name:    "many_items_with_small_weights_and_values",
			weights: []int{1, 2, 3},
			val:     []int{1, 2, 3},
			n:       3,
			W:       10,
			want:    10, // all have same value/weight ratio (1), so any combo summing to 10 gives 10
		},
		{
			name:    "best_ratio_in_middle_of_array",
			weights: []int{10, 1, 5},
			val:     []int{10, 2, 7},
			// ratios: 1.0, 2.0, 1.4 → best is index 1 (val=2, wt=1)
			n:    3,
			W:    7,
			want: 14, // 7 copies of item 1 → 7 * 2 = 14
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unboundedKnapsackOptimized(tt.weights, tt.val, tt.n, tt.W)
			assert.Equal(t, tt.want, got)
		})
	}
}
