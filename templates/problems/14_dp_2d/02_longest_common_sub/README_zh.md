# 最长公共子序列

- **难度**: 中等
- **分类**: 二维动态规划
- **考点**: 动态规划, 字符串
- **链接**: [NeetCode](https://neetcode.io/problems/longest-common-subsequence) | [力扣 1143](https://leetcode.cn/problems/longest-common-subsequence/)

## 题目描述

给定两个字符串 `text1` 和 `text2`，返回这两个字符串的最长公共子序列的长度。如果不存在公共子序列，返回 0。一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

## 示例

**示例 1:**

```
输入: text1 = "abcde", text2 = "ace"
输出: 3
解释: 最长公共子序列是 "ace"，它的长度为 3。
```

**示例 2:**

```
输入: text1 = "abc", text2 = "abc"
输出: 3
解释: 最长公共子序列是 "abc"，它的长度为 3。
```

**示例 3:**

```
输入: text1 = "abc", text2 = "def"
输出: 0
解释: 两个字符串没有公共子序列，返回 0。
```

## 约束条件

- `1 <= text1.length, text2.length <= 1000`
- `text1` 和 `text2` 仅由小写英文字符组成。

## 函数签名

```go
func longestCommonSubsequence(text1 string, text2 string) int
```
