# Divisors of given number

find the numbers which leaves no reminder.

i.e n = 36

these are the numbers which leaves no reminder `[1, 2, 3, 4, 6, 9, 12, 18, 36]`

**brute-force:** loop from 1 to n (inclusive) and check if `n` is divisible by `i` without reminder and add `i` to reminders

**optimal approach:**

* loop from 1 to square root of n and check if `n` is divisible by `i` without reminder and add `i` to reminders
  * check if `n/i != i` and add `n/i` to reminders
