# Star Patterns

**Rectangular Start Pattern:**

```sh
* * * * *
* * * * *
* * * * *
* * * * *
```

variables: `n, i, j`

`n` = number of rows

`i` = current row number (outer loop)

`j` = current column number (inner loop)

Data table: `n = 5`


| n (number of rows) | i (current row) | j (number of stars per row) |
| -------------------- | ----------------- | :---------------------------: |
| 5                  | 0               |              5              |
| 5                  | 1               |              5              |
| 5                  | 2               |              5              |
| 5                  | 3               |              5              |
| 5                  | 4               |              5              |

compare `n` and `i` to find the pattern value in `j`

Since `j` is equal to `n` for each iteration, we can conclude with,`j < n`

Pseudo Code:

```go
for i := 0; i < n; i++ {
  for j := 0; j < n; j++ {
     // print *
  }

  //print new line
}
```

**Right Angled Triangle:**

```sh
*
* *
* * *
* * * *
* * * * *
```

variables: `n, i, j`

`n` = number of rows

`i` = current row number (outer loop)

`j` = current column number (inner loop)

Data table: `n = 5`


| n (number of rows) | i (current row) | j (number of stars per row) |
| -------------------- | ----------------- | :---------------------------: |
| 5                  | 0               |              1              |
| 5                  | 1               |              2              |
| 5                  | 2               |              3              |
| 5                  | 3               |              4              |
| 5                  | 4               |              5              |

compare `n` and `i` to find the pattern value in `j`

since `j` grows each iteration by `i+1` we can conclude `j`with,`j < i + 1`

```go
for i := 0; i < n; i++ {
  for j := 0; j < i + 1; j++ {
     // print *
  }

  //print new line
}
```

**Right Angled Numbered Pyramid - I & II:**

```sh
# Right Angled Numbered Pyramid - I
1
1 2
1 2 3
1 2 3 4
1 2 3 4 5

# Right Angled Numbered Pyramid - II
1
2 2
3 4 3
4 4 4 4
5 5 5 5 5
```

variables: `n, i, j`

`n` = number of rows

`i` = current row number (outer loop)

`j` = current column number (inner loop)

Data table: **Right Angled Numbered Pyramid - I** `n = 5`


| n (number of rows) | i (current row) | j (numbers per row) | %d        |
| -------------------- | ----------------- | :-------------------: | ----------- |
| 5                  | 0               |          1          | 1         |
| 5                  | 1               |          2          | 1 2       |
| 5                  | 2               |          3          | 1 2 3     |
| 5                  | 3               |          4          | 1 2 3 4   |
| 5                  | 4               |          5          | 1 2 3 4 5 |

since `j` grows each iteration by 1 we can conclude `j` with, `j < i + 1`

`%d = j + 1`

Data table: **Right Angled Numbered Pyramid - II** `n = 5`


| n (number of rows) | i (current row) | j (numbers per row) | %d        |
| -------------------- | ----------------- | :-------------------: | ----------- |
| 5                  | 0               |          1          | 1         |
| 5                  | 1               |          2          | 2 2       |
| 5                  | 2               |          3          | 3 3 3     |
| 5                  | 3               |          4          | 4 4 4 4   |
| 5                  | 4               |          5          | 5 5 5 5 5 |

since `j` grows each iteration by 1 we can conclude `j` with, `j < i + 1`

`%d = i + 1`

**Inverted Right Pyramid**

```bash
* * * * *
* * * *
* * *
* *
*
```

`n = 5`


| n | i | j |
| --- | --- | :-- |
| 5 | 0 | 5 |
| 5 | 1 | 4 |
| 5 | 2 | 3 |
| 5 | 3 | 2 |
| 5 | 4 | 1 |

Compare `n` and `i` to find the pattern value in `j`

`j` differs `n- i` times each on iteration

Formula: `n - i`

**Inverted Right Numbered Pyramid**

```bash
1 2 3 4 5
1 2 3 4
1 2 3
1 2
1
```


| n | i | j |
| --- | --- | :-- |
| 5 | 0 | 5 |
| 5 | 1 | 4 |
| 5 | 2 | 3 |
| 5 | 3 | 2 |
| 5 | 4 | 1 |

Compare `n` and `i` to find the pattern value in `j`

`j` differs `n- i` times each on iteration

Formula: `n - i`

`%d = j + 1`

**Pyramid Pattern:**

```bash
*
      * * * 
    * * * * *
  * * * * * * *
* * * * * * * * *
```

`n = 5`


| n | i | s (space count) | j |
| --- | --- | :---------------: | --- |
| 5 | 0 |        4        | 1 |
| 5 | 1 |        3        | 3 |
| 5 | 2 |        2        | 5 |
| 5 | 3 |        1        | 7 |
| 5 | 4 |        0        | 9 |

Formula:
`s = n - i - 1`

`j = i * 2 + 1`

**Inverted Star Pyramid:**

```bash
* * * * * * * * *
  * * * * * * *
    * * * * *
      * * *
        *
```

`n = 5`


| n | i | s (space count) | j |
| --- | --- | :---------------: | --- |
| 5 | 0 |        0        | 9 |
| 5 | 1 |        1        | 7 |
| 5 | 2 |        2        | 5 |
| 5 | 3 |        3        | 3 |
| 5 | 4 |        4        | 1 |

Formula:
`s = i`
`j =  2 * n - (2 * i + 1)`

`j = 2(n - i) + 1`

2 * n - (2 * i + 1) = 2 * 5 - (2 * 0 + 1) = 10 - 1 = 9

2 * n - (2 * i + 1) = 2 * 5 - (2 * 1 + 1) = 10 - 3 = 7

2 * n - (2 * i + 1) = 2 * 5 - (2 * 2 + 1) = 10 - 5 = 5

2 * n - (2 * i + 1) = 2 * 5 - (2 * 3 + 1) = 10 - 7 = 3

2 * n - (2 * i + 1) = 2 * 5 - (2 * 4 + 1) = 10 - 9 = 1

**Diamond Star Pattern:**

```bash
*           = -3
      * * *         = -2
    * * * * *       = -1
  * * * * * * *     = 0
    * * * * *       = 1
      * * *         = 2
        *           = 3
```


| l | n | i | s | j |
| --- | :-: | :--: | :-: | --- |
| 3 | 3 | -3 | 3 | 1 |
| 2 | 3 | -2 | 2 | 3 |
| 1 | 3 | -1 | 1 | 5 |
| 0 | 3 | 0 | 0 | 7 |
| 1 | 3 | 1 | 1 | 5 |
| 2 | 3 | 2 | 2 | 3 |
| 3 | 3 | 3 | 3 | 1 |

*Core Logic for `l`:*

```go
if i < 0 {
  l = -i
} else {
  l = i
}
```

Notes: upon multiplying `l` with `i` based on the above condition, we get the `l` values as mentioned in the table above.

`s = l` since `l` and `s` is same in the table, we conclude with `s < l`

`j = 2(n - l) + 1` since `n - l + 1` can give right answer for `index -3`, it wont work for further iterations. So we multiply with `2` as `2(n-l)`and add `1`, we conclude with `j < 2(n-l) + 1`

**Half Diamond Star Pattern:**

```bash
*     = -2
**    = -1
***   = 0
**    = 1
*     = 2
```


| l | n | i  | j |
| --- | --- | ---- | --- |
| 2 | 2 | -2 | 1 |
| 1 | 2 | -1 | 2 |
| 0 | 2 | 0  | 3 |
| 1 | 2 | 1  | 2 |
| 2 | 2 | 2  | 1 |

```go
if i < 0 {
 l = -i 
} else {
 l = i
}
```

`j = n - l + 1`, If we perform `n - l + 1` we get the expected values in `j`, hence we conclude with `j < n - l + 1`

**Binary Number Triangle:**

```bash
1
0 1
1 0 1
0 1 0 1
1 0 1 0 1
```


| n | i | j |
| --- | --- | --- |
| 5 | 0 | 1 |
| 5 | 1 | 2 |
| 5 | 2 | 3 |
| 5 | 3 | 4 |
| 5 | 4 | 5 |

`j = i + 1`, for each iteration `j` has `i + 1` difference, hence we conclude with `j < i + 1`

💡 Since we need to print the binary numbers (0 and 1) we have to come up with conditional logic to print them.

💡 Since the odd rows always start with `1`, we can make the below condition

```go
if i%2 == 0 {
 start = 1
} else {
 start = 0 
}
```

🤔 Each time we print the number, we have to toggle between (`0` and `1`) this can be achieved by `start = 1 - start` **after printing the current start value**.

**Number Crown Pattern:**

```bash
1    1
12  21
123321

```


| n | i | j | s (space between) | k |
| --- | --- | --- | ------------------- | --- |
| 3 | 0 | 1 | 4                 | 1 |
| 3 | 1 | 2 | 2                 | 2 |
| 3 | 2 | 3 | 0                 | 3 |

`j = i + 1`

`s = 2 * (n-i-1)`

`k = i + 1`

end



[Youtube Tutorial](https://www.youtube.com/watch?v=cBCUXw7dC9c&t=55s)
