# 替换后的最长重复字符

- **难度**：中等
- **分类**：滑动窗口
- **考点**：字符串、哈希表、滑动窗口
- **链接**：[NeetCode](https://neetcode.io/problems/longest-repeating-substring-with-replacement) | [力扣 424](https://leetcode.cn/problems/longest-repeating-character-replacement/)

## 题目描述

给你一个字符串 `s` 和一个整数 `k`。你可以选择字符串中的任意一个字符，并将其更改为任何其他大写英文字母。该操作最多可执行 `k` 次。

在执行上述操作后，返回包含相同字母的最长子串的长度。

## 示例

**示例 1：**

```
输入：s = "ABAB", k = 2
输出：4
解释：用两个 'B' 替换为两个 'A'，反之亦然。整个字符串变为 "AAAA" 或 "BBBB"，长度为 4。
```

**示例 2：**

```
输入：s = "AABABBA", k = 1
输出：4
解释：将下标为 3 的 'B' 替换为 'A'，子串 "AABA" 的长度为 4，且所有字符相同。
```

**示例 3：**

```
输入：s = "AAAA", k = 0
输出：4
解释：无需替换。整个字符串已经由相同字符组成。
```

## 约束条件

- `1 <= s.length <= 10^5`
- `s` 仅由大写英文字母组成。
- `0 <= k <= s.length`

## 函数签名

```go
func characterReplacement(s string, k int) int
```
