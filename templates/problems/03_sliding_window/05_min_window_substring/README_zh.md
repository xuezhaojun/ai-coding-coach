# 最小覆盖子串

- **难度**：困难
- **分类**：滑动窗口
- **考点**：字符串、哈希表、滑动窗口
- **链接**：[NeetCode](https://neetcode.io/problems/minimum-window-with-characters) | [力扣 76](https://leetcode.cn/problems/minimum-window-substring/)

## 题目描述

给你一个字符串 `s`、一个字符串 `t`。返回 `s` 中涵盖 `t` 所有字符的最小子串。如果 `s` 中不存在涵盖 `t` 所有字符的子串，则返回空字符串 `""`。

注意：`t` 中的重复字符也必须被包含在窗口中，且数量不能少于 `t` 中该字符的数量。测试用例保证答案是唯一的。子串是字符串中连续的字符序列。

## 示例

**示例 1：**

```
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含了字符串 t 中的 'A'、'B' 和 'C'。
```

**示例 2：**

```
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 就是最小覆盖子串。
```

**示例 3：**

```
输入：s = "a", t = "aa"
输出：""
解释：t 中有两个 'a'，但 s 中只有一个 'a'，因此不存在有效的覆盖子串。
```

## 约束条件

- `m == s.length`
- `n == t.length`
- `1 <= m, n <= 10^5`
- `s` 和 `t` 由大写和小写英文字母组成。

## 函数签名

```go
func minWindow(s string, t string) string
```
