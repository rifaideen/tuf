# Booyer Moore's Voting Algorithm

Booyer Moore's algorithm used to find the majority element (n/2 times at-least)

start with `count = 0` and `element = none`

* start the loop `i` from `0` to `n - 1`
  * check if `count == 0`, then assign nums[i] to element and increment the counter
    * i.e `element = nums[i]` and `count++`
  * check else-if `element == nums[i]`, then increment the counter
    * i.e `count++`
  * else, decrement the counter
    * i.e count--

**pseudo-code:**

```bash
count = 0
element = -infinity

for i = 0; i < n - 1; i++ {
  if count == 0 {
     element = nums[i]
     count++
  } else if element == nums[i] {
    count++
  } else {
    count--
  }
}

return element
```
