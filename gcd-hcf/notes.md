# GCD & HCF of two numbers

Given two integers `num1` and `num2`,

while `num1` and `num2 ` > `0`

- if `num1 > num2` take modulo of `num1` by `num2`and assign it to `num1`
- else take modulo of `num2` by `num1` and assign it to `num2`

if `num1` becomes `0` , then `num2` becomes GCD, else `num1` becomes GCD

**pseudo-code:**

```bash
while num1 > 0 && num2 > 0 {
  if num1 > num2 {
    num1 = num1 % num2
  } else {
    num2 = num2 % num1
  }
}

return num1 == 0 ? num2 : num1
```
