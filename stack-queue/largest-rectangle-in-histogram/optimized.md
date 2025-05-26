Instead of calculating pse and nse before solving this problem, we will be calculating on the fly.

* **start the loop** from 0 to n-1
  * **calculate pse** for the current value

    * while popping the greater value from the stack, we**calculate the nse**
    * since the stack top is being popped up, whatever is the**current value in the outer loop is the nse** of the popped element
    * if the stack is non empty, whatever is on the**stack top is the pse** of the popped element in the inner loop,**or** we put**-1** (typical motonic stack indexing strategy)
    * once nse, pse and current value is set, calculate the max rectangle by applying the formula`(right - left - 1) * current`
  * push the current index to the stack
* if stack has something left out after this main loop
  * we calculate the nse, pse and current value
    * current = stack.top and pop the stack
    * nse will be n here
    * pse = -1 when stack becomes empty after popping, else stack top
  * apply the formula
