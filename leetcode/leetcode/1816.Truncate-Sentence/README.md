# [1816. Truncate Sentence](https://leetcode.com/problems/truncate-sentence/)

## 题目

A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each of the words consists of only uppercase and lowercase English letters (no punctuation).

- For example, "Hello World", "HELLO", and "hello world hello world" are all sentences.

You are given a sentence s and an integer k. You want to truncate s such that it contains only the first k words. Return s after truncating it.

**Example 1**:

    Input: s = "Hello how are you Contestant", k = 4
    Output: "Hello how are you"
    Explanation:
    The words in s are ["Hello", "how" "are", "you", "Contestant"].
    The first 4 words are ["Hello", "how", "are", "you"].
    Hence, you should return "Hello how are you".

**Example 2**:

    Input: s = "What is the solution to this problem", k = 4
    Output: "What is the solution"
    Explanation:
    The words in s are ["What", "is" "the", "solution", "to", "this", "problem"].
    The first 4 words are ["What", "is", "the", "solution"].
    Hence, you should return "What is the solution".

**Example 3**:

    Input: s = "chopper is not a tanuki", k = 5
    Output: "chopper is not a tanuki"

**Constraints:**

- 1 <= s.length <= 500
- k is in the range [1, the number of words in s].
- s consist of only lowercase and uppercase English letters and spaces.
- The words in s are separated by a single space.
- There are no leading or trailing spaces.

## 题目大意

句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。每个单词仅由大小写英文字母组成（不含标点符号）。

- 例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。

给你一个句子 s 和一个整数 k ，请你将 s 截断使截断后的句子仅含前 k 个单词。返回截断 s 后得到的句子。

## 解题思路

- 遍历字符串 s，找到最后一个空格的下标 end
- 如果 end 为 0，直接返回 s,否则返回 s[:end]

## 代码

```go
package leetcode

func truncateSentence(s string, k int) string {
	end := 0
	for i := range s {
		if k > 0 && s[i] == ' ' {
			k--
		}
		if k == 0 {
			end = i
			break
		}
	}
	if end == 0 {
		return s
	}
	return s[:end]
}
```