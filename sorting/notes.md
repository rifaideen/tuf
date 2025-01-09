# Sorting Algorithms

**Selection Sort:** [Tamil](https://www.youtube.com/watch?v=V6Gpt1IOCUI) | [English](https://www.youtube.com/watch?t=167&v=HGk_ypEuS24&feature=youtu.be)

upon each iteration, lowest element inserted first (sorts in left to right direction)

---

* In selection start, we start outer loop `i` from `0`to`n-1` (i.e i < n-1)
  * start inner loop `j` from `i+1` to `n` (i.e j < n)
    * compare value of `i` and `j` and swap when values `i > j`
    * after iteration 0, `lowest element inserted first` index
    * after iteration 1, `second lowest element inserted second` index, and so on...

```bash
nums = [5, 4, 3, 2, 1]
        0  -  -  -  -  after iteration 0,  min at 0
        0  1  -  -  -  after iteration 1,  min at 1
        0  1  2  -  -  after iteration 2,  min at 2
        0  1  2  3  -  after iteration 3,  min at 3
```

**Bubble Sort:** [Tamil](https://www.youtube.com/watch?v=_PNIffgsSNI) | [English](https://www.youtube.com/watch?v=HGk_ypEuS24&t=1061s)

upon each iteration, largest elements are inserted last (sorted in right to left direction)

---

* In bubble sort, we start **outer loop** `i` from `0`to`n-1` (i.e i < len(nums)-1)
  * start **inner loop** `j` from `0`to`n - i` (i.e j < (len(nums) - 1) - i)
    * after first iteration, `largest number` is stored in `last index`
    * after second iteration, second largest number is stored in `second last index`, and so on.

```bash
nums = [5, 4, 3, 2, 1]
        -  -  -  -  4   after iteration 0, max at 4
        -  -  -  3  4   after iteration 1, max at 3
        -  -  2  3  4   after iteration 2, max at 2
        -  1  2  3  4   after iteration 3, max at 1
```

**pseudo-code:**

```bash
for i = n-1; i >= 1; i++ {
  for j = 0; j <= i-1; j++ {
    if nums[j] > num[j+1] {
       nums[j+1], nums[j] = nums[j], nums[j+1]
    }
  }
}
```

**Insertion Sort:** [Tamil](https://www.youtube.com/watch?v=GsWyL9RdE3A) | [English](https://www.youtube.com/watch?v=HGk_ypEuS24&t=1900s)

splits the array into `sorted (left)` and `non-sorted(right)` and insert the non-sorted element within sorted elements (left) at correct position.

---

* start the outer loop `i` from `1` to `n-1`
  * store `k = nums[i]`
  * start the inner loop `j` from `i-1` to `0` and `nums[j] > k`
    * move value of `j` to `j+1`
  * add `k` to `j+1`

Visual Representation:

```bash
nums = [5, 4, 3, 2, 1]
        0  -  -  -  -  iteration 0, 1 element in sorted,  4 elements in un-sorted
        0  1  -  -  -  iteration 1, 2 elements in sorted, 3 elements in un-sorted
        0  1  2  -  -  iteration 2, 3 elements in sorted, 2 elements in un-sorted
        0  1  2  3  -  iteration 3, 4 elements in sorted, 1 element  in un-sorted

+ sorted
- unsorted
```
