# [1022. Sum of Root To Leaf Binary Numbers](https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/)

## 题目

You are given the root of a binary tree where each node has a value 0 or 1. Each root-to-leaf path represents a binary number starting with the most significant bit.

For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.
For all leaves in the tree, consider the numbers represented by the path from the root to that leaf. Return the sum of these numbers.

The test cases are generated so that the answer fits in a 32-bits integer.

**Example 1:**

```c
Input: root = [1,0,1,0,1,0,1]
Output: 22
Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
```

**Example 2:**

```c
Input: root = [0]
Output: 0
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 1000]`.

- `Node.val` is `0` or `1`.


## 题目大意

给定一棵结点值都是`0`或`1`的二叉树，每条从根结点到叶结点的路径都代表一个从最高有效位开始的二进制数。

返回从根节点到所有叶结点的路径所表示的数字之和。


## 解题思路

采用递归的方式对根结点`root`进行后序遍历（左子树-右子树-根结点）。

**递归函数的返回值**：

递归遍历每个结点时，计算从根结点到当前访问结点的所表示数值`sum`都用到了上次的计算结果，所以递归函数的返回值是当前访问结点的计算结果值。

**递归函数的逻辑**：

- 当前遍历结点为`nil`，表示本层递归结束了，直接`return 0`。

- 如果当前访问结点是叶结点，则返回从根结点到该结点所表示的数值`sum`。
- 如果当前访问结点不是叶结点，则返回左子树和右子树所对应的结果之和。
