# Reverse Number

While `n > 0` perform the below steps

1. find the reminder by performing modulo operation  `reminder = n % 10`
2. multiply the reversed number by 10 and add the reminder `reversed = (reversed * 10) + reminder `
3. update the value of `n` by 10 `n = n/10`

Pseudo-code:

```bash
reversed = 0

while n > 0 {
 // find the reminder by performing the modulo of 10
 reminder = n % 10

 // update the reversed value
 reversed = (reversed * 10) + reminder

 // update the n value
 n = n / 10
}

return reversed
```
