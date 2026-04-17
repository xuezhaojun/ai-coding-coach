# 无重复字符的最长子串

- **难度**：中等
- **分类**：滑动窗口
- **考点**：字符串、哈希表、滑动窗口
- **链接**：[NeetCode](https://neetcode.io/problems/longest-substring-without-duplicates) | [力扣 3](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

## 题目描述

给定一个字符串 `s`，请你找出其中不含有重复字符的最长子串的长度。

子串是字符串中连续的非空字符序列。子串中不能包含任何重复字符。

## 示例

**示例 1：**

```
输入：s = "abcabcbb"
输出：3
解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

**示例 2：**

```
输入：s = "bbbbb"
输出：1
解释：因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

**示例 3：**

```
输入：s = "pwwkew"
输出：3
解释：因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意 "pwke" 是子序列而不是子串。
```

## 约束条件

- `0 <= s.length <= 5 * 10^4`
- `s` 由英文字母、数字、符号和空格组成。

## 函数签名

```go
func lengthOfLongestSubstring(s string) int
```
