## Infix to Postfix Conversion

to convert the given infix expression to postfix expression we use stack, ans and priority mapping

Assign priorities to various operators and rest of them should be given -1 priority

```
priority = {
 "^": 3,
 "*": 2,
 "/": 2,
 "+" 1,
 "-": 1
}
```

start the loop and check if the charecter is operand or operand or open bracket or close bracket.

* push the operands to the answer
* push the open bracket`(` to the stack
* for close bracket`(`, pop the values from the stack until the open bracket`(` found and push it to answer
* push the operator to the stack
  * if the current operator priority is less than or equal to stack top's priority, pop the top element and push it to answer
  * push it to the stack

once, the iterations are over, pop the stack and push it to answers
