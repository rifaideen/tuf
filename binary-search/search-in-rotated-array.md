**Search in rotated sorted array - I**

using binary search, we can solve this. But we cannot directly eliminate the left half or right half, since the array is sorted rotated.

**Approach:**
We need to `identify (sorted side)` which side of the array is sorted. Left or Right?

left side is rotated, when `nums[left] <= nums[mid]` else, right side is rotated for sure

Once identified, check if the target is within the range.

**i.e: for left sorted:** The target should be within left and mid.

`if nums[left] <= target && target <= nums[mid]` right = mid - 1. `else` left = mid + 1

**i.e: for right sorted:** The target should be within mid and right

`if nums[mid] <= target && target <= nums[right]` left = mid + 1. `else` hight = mid - 1

**Search in rotated sorted array - II**

everything is same except when left == mid == right, we can decide which side to go. In that case, we shrink the seach space on both side and continue.

i.e:

```go
if nums[left] == nums[mid] && nums[mid] == nums[right] {
    left++
    right--
    continue
}
```
