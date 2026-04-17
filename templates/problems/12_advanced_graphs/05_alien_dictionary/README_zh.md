# 火星词典

- **难度**: 困难
- **分类**: 高级图
- **考点**: 拓扑排序, 广度优先搜索, 图, 字符串
- **链接**: [NeetCode](https://neetcode.io/problems/foreign-dictionary) | [力扣 269](https://leetcode.cn/problems/alien-dictionary/) | [LintCode 892](https://www.lintcode.com/problem/892/)

## 题目描述

现有一种使用英语字母的外星语言，但字母的顺序未知。

给你一个字符串列表 `words`，该列表来自外星语言的字典，其中字符串已经按照该新语言的字典序排列。

请你推导出该语言中字母的顺序。如果顺序无效（即存在环或矛盾），返回空字符串 `""`。如果有多个有效顺序，返回其中任意一个。

## 示例

**示例 1:**

```
输入: words = ["wrt","wrf","er","ett","rftt"]
输出: "wertf"
```

**示例 2:**

```
输入: words = ["z","x"]
输出: "zx"
```

**示例 3:**

```
输入: words = ["z","x","z"]
输出: ""
解释: 顺序无效，因为 'z' 必须同时排在 'x' 之前和之后。
```

## 约束条件

- `1 <= words.length <= 100`
- `1 <= words[i].length <= 100`
- `words[i]` 仅由小写英文字母组成。

## 函数签名

```go
func alienOrder(words []string) string
```
