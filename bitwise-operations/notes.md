# Bitwise Operations:

**Swap two numbers**

to swap two numbers without using temp variable, we use bit manipulation with xor operator.

i.e: a = 5, b = 6

expected: a = 6, b = 5

* perform xor operator on a with b`a = a^b`
* perform xor operator on b with a`b = a^b`
* perform xor operator again on a with b`a = a^b`

i.e

```go
a = a^b
// at this point, a contains a^b, so a^b is equivalent to (a^b) ^ b
// b ^ b's are cancelled hence it has a
b = a^b
// at this point, b contains a, so a^b is equivalent to (a^b) ^ b, since b contains a it is equivalent to (a^b) ^ b
// a ^ a's are cancelled hence it is b
a = a^b
```

**Check if i'th bit is set or not**

to check if i'th bit is set or not, we need to shift the 1 to the left or right by given number of times and perform the logical and (&) operation. n = given number, i = bit position (0 based from right side)

- using left shift
  - (n & (1 << i)) > 0 ? true : false
- using right shift
  - ((n >> i) & 1) > 0 ? true : false

**Set the i-th bit**

to set the i-th bit, we need to left shift 1 by i times and perform logical or (|) with the given number. n = given number, i = bit position

* n | (1 << i)


**Clear the i-th bit**

to clear the i-th bit, we need to left shift 1 by i and negate (~) the result and perform and operation with n

* n & (~(1 << i))


**Toggle the i-th bit**

to toggle the i-th bit, we need to left shift 1 by i and perform xor operation (^) with n

* n ^ (1 << i)
