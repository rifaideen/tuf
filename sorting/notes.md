# Sorting

**Bubble Sort:** [Tamil Tutorial](https://www.youtube.com/watch?v=_PNIffgsSNI) [English Tutorial](https://www.youtube.com/watch?v=HGk_ypEuS24&t=1061s)

* In bubble sort, we **start outer loop** from `n-1` to `1` (inclusive)
* start inner loop from `0` to `i-1`
* after first iteration, `largest number` is stored in `last index`
* after second iteration, second largest number is stored in `second last index`, and so on.

**input:** `nums = [5,4,3,2,1]`

**output:** `[1,2,3,4,5]`

```bash
for i = n-1; i >= 1; i++ {
  for j = 0; j <= i-1; j++ {
    if nums[j] > num[j+1] {
       nums[j+1], nums[j] = nums[j], nums[j+1]
    }
  }
}
```

* Iteration 0: `5, 4, 3, 2, 1` i = 4, j = 0 to 3

  * `nums[0] > nums[1] = true = swap(5, 4) = 4, 5, 3, 2, 1`
  * `nums[1] > nums[2] = true = swap(5, 3) = 4, 3, 5, 2, 1`
  * `nums[2] > nums[3] = true = swap(5, 2) = 4, 3, 2, 5, 1`
  * `nums[3] > nums[4] = true = swap(5, 1) = 4, 3, 2, 1, 5`
* Iteration 1: `4, 3, 2, 1, 5` i = 3, j = 0 to 2

  * `nums[0] > nums[1] = true = swap(4, 3) = 3, 4, 2, 1, 5`
  * `nums[1] > nums[2] = true = swap(4, 2) = 3, 2, 4, 1, 5`
  * `nums[2] > nums[3] = true = swap(4, 1) = 3, 2, 1, 4, 5`
* Iteration 2: `3, 2, 1, 4, 5` i = 2, j = 0, 1

  * `nums[0] > nums[1] = true = swap(3, 2) = 2, 3, 1, 4, 5`
  * `nums[1] > nums[2] = true = swap(3, 1) = 2, 1, 3, 4, 5`
* Iteration 3: `2, 1, 3, 4, 5` i = 1, j = 0 to 0

  * `nums[0] > nums[1] = true = swap(2, 1) = 1, 2, 3, 4, 5`
