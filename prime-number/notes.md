# Prime Number Check

prime number is something, which is divisible by `1` and the number `itself`.


**brute-force:** `O(n)` time complexity

* loop through `i` from `1` to `n` and check if`n % i == 0` the increment the counter.
* check if counter = 2, then `true` else `false`

**pseudo-code:**

```bash
counter = 0

for i = 1; i <= n; i++ {
  if n % i == 0 {
    counter++
  }
}

return counter == 2
```

**Optimized solution:** `O(sqrt(n))` time complexity

* loop through `i` from `1` to `sqrt(n)`, check if `n % i == 0` and increment the counter when true
  * check if `n / i != i`, then increment the counter
* check if counter = 2, then `true` else `false`

**pseudo-code:**

```bash
counter = 0

for i = 1; i <= sqrt(n); i++ {
  if n % i == 0 {
    counter++
  
    if n / i != i {
      counter++
    }
  }
}

return counter == 2
```
