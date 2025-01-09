# Sorting Algorithms

**Selection Sort:** [Tamil](https://www.youtube.com/watch?v=V6Gpt1IOCUI) [English](https://www.youtube.com/watch?t=167&v=HGk_ypEuS24&feature=youtu.be)

upon each iteration, lowest element inserted first

---


* In selection start, we start outer loop `i` from `0`to`n-1` (i.e i < n-1)
  * start inner loop `j` from `i+1` to `n` (i.e j < n)
    * compare value of `i` and `j` and swap when values `i > j`
    * after iteration 0, `lowest element inserted first` index
    * after iteration 1, `second lowest element inserted second` index, and so on...

**Bubble Sort:** [Tamil](https://www.youtube.com/watch?v=_PNIffgsSNI) | [English](https://www.youtube.com/watch?v=HGk_ypEuS24&t=1061s)

upon each iteration, largest elements are inserted last

---



* In bubble sort, we start **outer loop** `i` from `0`to`n-1` (i.e i < len(nums)-1)
  * start **inner loop** `j` from `0`to`n - i` (i.e j < (len(nums) - 1) - i)
    * after first iteration, `largest number` is stored in `last index`
    * after second iteration, second largest number is stored in `second last index`, and so on.

**input:** `nums = [5,4,3,2,1]`

**output:** `[1,2,3,4,5]`

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
