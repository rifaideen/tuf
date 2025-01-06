# Armstrong Number

armstrong number is something which is equal to the sum of individual digit raised to the power of total number of digits.

i.e 153

`1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153`

Pseudo-code:

```bash
original = n #given number
sum = 0 #start with 0
pow = math.log10(n) + 1 #number of digits in n

while n > 0 {
  reminder = n % 10 #find modulo reminder
  sum += n^pow #sum of existing value and n^pow
  n /= 10
}

return sum == original #check if sum is equal to original
```

⛔️ do not rush into converting the given integer to string and count. which would require more space and less efficient.
