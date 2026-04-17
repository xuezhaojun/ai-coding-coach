# 字符串的排列

- **难度**：中等
- **分类**：滑动窗口
- **考点**：字符串、哈希表、滑动窗口、双指针
- **链接**：[NeetCode](https://neetcode.io/problems/permutation-string) | [力扣 567](https://leetcode.cn/problems/permutation-in-string/)

## 题目描述

给你两个字符串 `s1` 和 `s2`，如果 `s2` 包含 `s1` 的排列，则返回 `true`；否则，返回 `false`。

换句话说，如果 `s1` 的某个排列是 `s2` 的子串，则返回 `true`。排列是指对字符串中所有字符的重新排列。

## 示例

**示例 1：**

```
输入：s1 = "ab", s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 "ba"，它是 s2 的子串。
```

**示例 2：**

```
输入：s1 = "ab", s2 = "eidboaoo"
输出：false
解释："ab" 的任何排列都不是 s2 中的连续子串。
```

**示例 3：**

```
输入：s1 = "aab", s2 = "ccccbaa"
输出：true
解释：s2 包含 s1 的排列 "baa" 作为子串。
```

## 约束条件

- `1 <= s1.length, s2.length <= 10^4`
- `s1` 和 `s2` 仅包含小写英文字母。

## 函数签名

```go
func checkInclusion(s1 string, s2 string) bool
```
