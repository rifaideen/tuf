# Pre, In and Post order traversal in one traversal

We solve using a stack which contains node and value. We push the root initially and start the loop.

Create 3 arrays to store the result of pre, post and in-order lists

**Rules:**

* pop the element and follow the steps:
  * `value = 1`, push it to pre-order list, increment the value and re-insert, if left node present, insert left node with value 1
  * `value = 2`, push it to in-order list, increment the value and re-insert, if right node present,
  * insert right node with value 1
  * `value = 1`, push it to post-order list and node will be removed from stack
