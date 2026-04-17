# 分割回文串

- **难度**: 中等
- **分类**: 回溯
- **考点**: 回溯, 字符串, 回文, 动态规划
- **链接**: [NeetCode](https://neetcode.io/problems/palindrome-partitioning) | [力扣 131](https://leetcode.cn/problems/palindrome-partitioning/)

## 题目描述

给你一个字符串 `s`，请你将 `s` 分割成一些子串，使每个子串都是回文串。返回 `s` 所有可能的分割方案。

回文串是正着读和反着读都一样的字符串。每种分割方案必须恰好使用字符串中的每个字符一次。

## 示例

**示例 1:**

```
输入: s = "aab"
输出: [["a","a","b"], ["aa","b"]]
```

**示例 2:**

```
输入: s = "a"
输出: [["a"]]
```

**示例 3:**

```
输入: s = "aaa"
输出: [["a","a","a"], ["a","aa"], ["aa","a"], ["aaa"]]
```

## 约束条件

- `1 <= s.length <= 16`
- `s` 仅由小写英文字母组成。

## 函数签名

```go
func partition(s string) [][]string
```
