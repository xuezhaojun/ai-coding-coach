# 划分字母区间

- **难度**: 中等
- **分类**: 贪心
- **考点**: 字符串, 贪心, 哈希表, 双指针
- **链接**: [NeetCode](https://neetcode.io/problems/partition-labels) | [力扣 763](https://leetcode.cn/problems/partition-labels/)

## 题目描述

给你一个字符串 `s`。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 `s`。返回一个表示每个字符串片段长度的列表。

## 示例

**示例 1:**

```
输入: s = "ababcbacadefegdehijhklij"
输出: [9,7,8]
解释: 划分结果为 "ababcbaca"、"defegde"、"hijhklij"。每个字母最多出现在一个片段中。像 "ababcbacadefegde"、"hijhklij" 这样的划分是错误的，因为划分的片段数较少。
```

**示例 2:**

```
输入: s = "eccbbbbdec"
输出: [10]
解释: 所有字符都交织在一起，因此整个字符串是一个片段。
```

**示例 3:**

```
输入: s = "abcdef"
输出: [1,1,1,1,1,1]
解释: 每个字符只出现一次，所以每个字符单独成为一个片段。
```

## 约束条件

- `1 <= s.length <= 500`
- `s` 仅由小写英文字母组成。

## 函数签名

```go
func partitionLabels(s string) []int
```
