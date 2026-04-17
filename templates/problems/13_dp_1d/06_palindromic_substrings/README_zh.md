# 回文子串

- **难度**: 中等
- **分类**: 一维动态规划
- **考点**: 动态规划, 字符串, 双指针
- **链接**: [NeetCode](https://neetcode.io/problems/palindromic-substrings) | [力扣 647](https://leetcode.cn/problems/palindromic-substrings/)

## 题目描述

给你一个字符串 `s`，请你统计并返回这个字符串中回文子串的数目。回文字符串是正着读和倒过来读一样的字符串。子字符串是字符串中的连续字符序列。注意，即使子串内容相同，但起始或结束位置不同，也视为不同的子串。

## 示例

**示例 1:**

```
输入: s = "abc"
输出: 3
解释: 三个回文子串为 "a"、"b"、"c"。
```

**示例 2:**

```
输入: s = "aaa"
输出: 6
解释: 六个回文子串为 "a"（出现 3 次）、"aa"（出现 2 次）、"aaa"（出现 1 次）。
```

**示例 3:**

```
输入: s = "racecar"
输出: 10
```

## 约束条件

- `1 <= s.length <= 1000`
- `s` 仅由小写英文字母组成。

## 函数签名

```go
func countSubstrings(s string) int
```
