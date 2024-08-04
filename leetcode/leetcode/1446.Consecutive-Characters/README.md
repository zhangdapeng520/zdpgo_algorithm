# [1446. Consecutive Characters](https://leetcode.com/problems/consecutive-characters/)

## 题目

The power of the string is the maximum length of a non-empty substring that contains only one unique character.

Given a string s, return the power of s.

**Example 1**:

    Input: s = "leetcode"
    Output: 2
    Explanation: The substring "ee" is of length 2 with the character 'e' only.

**Example 2**:

    Input: s = "abbcccddddeeeeedcba"
    Output: 5
    Explanation: The substring "eeeee" is of length 5 with the character 'e' only.

**Example 3**:

    Input: s = "triplepillooooow"
    Output: 5

**Example 4**:

    Input: s = "hooraaaaaaaaaaay"
    Output: 11

**Example 5**:

    Input: s = "tourist"
    Output: 1

**Constraints:**

- 1 <= s.length <= 500
- s consists of only lowercase English letters.

## 题目大意

给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。

请你返回字符串的能量。

## 解题思路

- 顺序遍历进行统计

## 代码

```go
package leetcode

func maxPower(s string) int {
	cur, cnt, ans := s[0], 1, 1
	for i := 1; i < len(s); i++ {
		if cur == s[i] {
			cnt++
		} else {
			if cnt > ans {
				ans = cnt
			}
			cur = s[i]
			cnt = 1
		}
	}
	if cnt > ans {
		ans = cnt
	}
	return ans
}

```