# 交错字符串

- **难度**: 中等
- **分类**: 二维动态规划
- **考点**: 动态规划, 字符串
- **链接**: [NeetCode](https://neetcode.io/problems/interleaving-string) | [力扣 97](https://leetcode.cn/problems/interleaving-string/)

## 题目描述

给定三个字符串 `s1`、`s2`、`s3`，请你判断 `s3` 是否是由 `s1` 和 `s2` 交错组成的。两个字符串 `s` 和 `t` 交错的定义是：将 `s` 和 `t` 分别拆分为若干子串，然后将这些子串交替拼接形成 `s3`，同时保持 `s` 和 `t` 中字符的相对顺序不变。

![interleaving-string](assets/interleave.jpg)

## 示例

**示例 1:**

```
输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
解释: s1 和 s2 的字符在保持各自相对顺序的前提下交替拼接形成了 s3。
```

**示例 2:**

```
输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false
```

**示例 3:**

```
输入: s1 = "", s2 = "", s3 = ""
输出: true
```

## 约束条件

- `0 <= s1.length, s2.length <= 100`
- `0 <= s3.length <= 200`
- `s1`、`s2` 和 `s3` 仅由小写英文字母组成。

## 函数签名

```go
func isInterleave(s1 string, s2 string, s3 string) bool
```
