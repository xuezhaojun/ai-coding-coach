# 最长回文子串

- **难度**: 中等
- **分类**: 一维动态规划
- **考点**: 动态规划, 字符串, 双指针
- **链接**: [NeetCode](https://neetcode.io/problems/longest-palindromic-substring) | [力扣 5](https://leetcode.cn/problems/longest-palindromic-substring/)

## 题目描述

给你一个字符串 `s`，找到 `s` 中最长的回文子串。回文串是正读和反读都相同的字符串。如果有多个长度相同的答案，返回其中任意一个即可。

## 示例

**示例 1:**

```
输入: s = "babad"
输出: "bab"
解释: "aba" 同样是有效答案。
```

**示例 2:**

```
输入: s = "cbbd"
输出: "bb"
```

**示例 3:**

```
输入: s = "racecar"
输出: "racecar"
解释: 整个字符串就是一个回文串。
```

## 约束条件

- `1 <= s.length <= 1000`
- `s` 仅由小写英文字母组成。

## 函数签名

```go
func longestPalindrome(s string) string
```
