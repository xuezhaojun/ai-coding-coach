# 有效的字母异位词

- **难度**：简单
- **分类**：数组与哈希
- **考点**：字符串、哈希表、排序
- **链接**：[NeetCode](https://neetcode.io/problems/is-anagram) | [力扣 242](https://leetcode.cn/problems/valid-anagram/)

## 题目描述

给定两个字符串 `s` 和 `t`，如果 `t` 是 `s` 的字母异位词，则返回 `true`，否则返回 `false`。

字母异位词是通过重新排列另一个单词的所有字母而形成的单词或短语，每个字母恰好使用一次。

## 示例

**示例 1：**

```
输入：s = "anagram", t = "nagaram"
输出：true
解释："nagaram" 是由 "anagram" 的所有字母重新排列而成。
```

**示例 2：**

```
输入：s = "rat", t = "car"
输出：false
解释："rat" 中的字母无法重新排列为 "car"。
```

**示例 3：**

```
输入：s = "", t = ""
输出：true
解释：两个空字符串是平凡的字母异位词。
```

## 约束条件

- `0 <= s.length, t.length <= 5 * 10^4`
- `s` 和 `t` 仅由小写英文字母组成。

## 函数签名

```go
func isAnagram(s string, t string) bool
```
