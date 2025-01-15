# Kadane's Algorithm to find maximum sub-array sum

kadane's algorithm is used to find the maximum sub-array sum in the given array.

we start with two variables `sum` to store overall sum and `maxSum` to store maximum sum we have seen so far.

i.e `sum = 0`, `maxSum = -infinity`

* start the loop `i` from `0` to `len(nums) - 1`

  * calculate sum `sum += nums[i]`
  * set the max sum `maxSum = max(maxSum, sum)`
  * check if `sum <= 0`, reset `sum = 0` when true
* return `maxSum`
