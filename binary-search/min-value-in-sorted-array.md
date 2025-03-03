# Find min value in sorted rotated array

* using binary search, we find the mid value
* identify which side is sorted
  * if left side is sorted, pickup the minimum of answer and nums[left]. i.e: `ans = min(ans, nums[left])`
    * picked up the min value from the left side, move the pointer to the right
      * `left = mid + 1`
  * else, right side is sorted, pickup the minimum of answer and nums[mid]. i.e: `ans = min(ans, nums[mid])`
    * picked up the min value from the right side, move the pointer to the left
      * `right = mid - 1`
