# [396. Rotate Function](https://leetcode.com/problems/rotate-function/)

## 题目

You are given an integer array `nums` of length `n`.

Assume `arrk` to be an array obtained by rotating `nums` by `k` positions clock-wise. We define the **rotation function** `F` on `nums` as follow:

- `F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]`.

Return the maximum value of `F(0), F(1), ..., F(n-1)`.

The test cases are generated so that the answer fits in a **32-bit** integer.

**Example 1:**

```c
Input: nums = [4,3,2,6]
Output: 26
Explanation:
F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
So the maximum value of F(0), F(1), F(2), F(3) is F(3) = 26.
```

**Example 2:**

```c
Input: nums = [100]
Output: 0
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 105`
- `-100 <= nums[i] <= 100`

## 题目大意

给定一个长度为`n`的整数数组`nums`，设`arrk`是数组`nums`顺时针旋转`k`个位置后的数组。

定义`nums`的旋转函数`F`为：

- `F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]`

返回`F(0), F(1), ..., F(n-1)`中的最大值。

## 解题思路

**抽象化观察：**

```c
nums = [A0, A1, A2, A3]

sum  = A0 + A1 + A2+ A3
F(0) = 0*A0 +0*A0 + 1*A1 + 2*A2 + 3*A3
    
F(1) = 0*A3 + 1*A0 + 2*A1 + 3*A2 
     = F(0) + (A0 + A1 + A2) - 3*A3 
     = F(0) + (sum-A3) - 3*A3 
     = F(0) + sum - 4*A3

F(2) = 0*A2 + 1*A3 + 2*A0 + 3*A1 
     = F(1) + A3 + A0 + A1 - 3*A2 
     = F(1) + sum - 4*A2

F(3) = 0*A1 + 1*A2 + 2*A3 + 3*A0 
     = F(2) + A2 + A3 + A0 - 3*A1 
     = F(2) + sum - 4*A1
    
// 记sum为nums数组中所有元素和
// 可以猜测当0 ≤ i < n时存在公式:
F(i) = F(i-1) + sum - n * A(n-i)
```

**数学归纳法证明迭代公式：**

根据题目中给定的旋转函数公式可得已知条件：

- `F(0) = 0×nums[0] + 1×nums[1] + ... + (n−1)×nums[n−1]`；

- `F(1) = 1×nums[0] + 2×nums[1] + ... + 0×nums[n-1]`。

令数组`nums`中所有元素和为`sum`，用数学归纳法验证：当`1 ≤ k < n`时，`F(k) = F(k-1) + sum - n×nums[n-k]`成立。

**归纳奠基**：证明`k=1`时命题成立。

```c
F(1) = 1×nums[0] + 2×nums[1] + ... + 0×nums[n-1]
     = F(0) + sum - n×nums[n-1]
```

**归纳假设**：假设`F(k) = F(k-1) + sum - n×nums[n-k]`成立。

**归纳递推**：由归纳假设推出`F(k+1) = F(k) + sum - n×nums[n-(k+1)]`成立，则假设的递推公式成立。

```c
F(k+1) = (k+1)×nums[0] + k×nums[1] + ... + 0×nums[n-1]
       = F(k) + sum - n×nums[n-(k+1)] 
```

因此可以得到递推公式：

- 当`n = 0`时，`F(0) = 0×nums[0] + 1×nums[1] + ... + (n−1)×nums[n−1]`
- 当`1 ≤ k < n`时，`F(k) = F(k-1) + sum - n×nums[n-k]`成立。

循环遍历`0 ≤ k < n`，计算出不同的`F(k)`并不断更新最大值，就能求出`F(0), F(1), ..., F(n-1)`中的最大值。
