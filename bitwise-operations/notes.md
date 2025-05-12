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

**Check if number is odd or not**

to check if the number is odd or not, we need to check if the right most bit is set or not by simply placing 1 underneath and perform logical and operation

* (n & 1) > 0

**Check if the number is power of 2 or not**

to check if the number is power of 2 or not, we need to perform logical and operation with the given number and -1 of given number should give us 0

* (n & (n-1)) == 0

n = 8

(1 0 0 0 & (0 1 1 1)) == 0

1 0 0 0 &

0 1 1 1

0 0 0 0 (final result)

**Count the number of set bits**

to count the number of set bits, we need to have counter variable and perform logic and with given number and 1 then we need to shift the given number to the right by 1 bit.

```bash
count = 0
while n > 0 {
  count += (n & 1) // when the right most bit is set, it will give us 1, otherwise 0
  n = n >> 1 // shift the n to the right by 1 bit
}
```

**Remove the last set bit**

to remove the last set bit, we need to perform the logical and operation with the given number and -1 of given number

* n & (n-1)

n = 12

1 1 0 0 (n = 12) &

1 0 1 1 (n - 1 = 11)

1 0 0 0
