# 1.Two Sum

**Input:**

Given array of `non-sorted` integers and `target` integer.

**Output:**

return `indices` of `two numbers` whose sum matches the `target`.

**Statement:**

Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to `target`* .

You may assume that each input would have ***exactly* one solution** , and **you may not use the *same* element twice**.

You can return the answer in any order.

**Approach:**

*create map to track visited values and it's index i.e `map[value] = index`. On each iteration find the value of `x` and see if `x` present in map.*

💡 Add the visited value and it's index to the map.

**Formula:**

`current value` + `x` = `target`, we can rewrite into `x`=`target`-`current value`

**Pseudo-Code:**

```go
// create visited map with value:index pair
visited := map[int]int{}

for i := 0; i < len(nums); i++ {
  current := nums[i]
  x := target - current

  // check if x present in visited map.
  if visited[x] {
   // x found, return the current and x indicex
   return []int{ visited[x], nums[current] }
  }
  
  // add the current value and index
  visited[current] = i
}

```

end
