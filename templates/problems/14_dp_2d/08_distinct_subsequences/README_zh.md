# 不同的子序列

- **难度**: 困难
- **分类**: 二维动态规划
- **考点**: 动态规划, 字符串
- **链接**: [NeetCode](https://neetcode.io/problems/distinct-subsequences) | [力扣 115](https://leetcode.cn/problems/distinct-subsequences/)

## 题目描述

给你两个字符串 `s` 和 `t`，统计并返回在 `s` 的子序列中 `t` 出现的个数。字符串的子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。题目数据保证答案符合 32 位带符号整数范围。

## 示例

**示例 1:**

```
输入: s = "rabbbit", t = "rabbit"
输出: 3
解释: 有三种方式可以从 "rabbbit" 中选取 "rabbit"（通过跳过三个 'b' 中的不同一个）。
```

**示例 2:**

```
输入: s = "babgbag", t = "bag"
输出: 5
解释: "babgbag" 中有 5 个不同的子序列等于 "bag"。
```

**示例 3:**

```
输入: s = "aaa", t = "a"
输出: 3
解释: 三个 'a' 字符中的每一个都构成一个有效子序列。
```

## 约束条件

- `1 <= s.length, t.length <= 1000`
- `s` 和 `t` 由英文字母组成。

## 函数签名

```go
func numDistinct(s string, t string) int
```
