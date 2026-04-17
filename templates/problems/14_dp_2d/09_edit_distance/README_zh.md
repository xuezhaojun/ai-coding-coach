# 编辑距离

- **难度**: 中等
- **分类**: 二维动态规划
- **考点**: 动态规划, 字符串
- **链接**: [NeetCode](https://neetcode.io/problems/edit-distance) | [力扣 72](https://leetcode.cn/problems/edit-distance/)

## 题目描述

给你两个单词 `word1` 和 `word2`，请返回将 `word1` 转换成 `word2` 所使用的最少操作数。你可以对一个单词进行如下三种操作：插入一个字符、删除一个字符、替换一个字符。

## 示例

**示例 1:**

```
输入: word1 = "horse", word2 = "ros"
输出: 3
解释: horse -> rorse（将 'h' 替换为 'r'）-> rose（删除 'r'）-> ros（删除 'e'）。
```

**示例 2:**

```
输入: word1 = "intention", word2 = "execution"
输出: 5
解释: intention -> inention（删除 't'）-> enention（将 'i' 替换为 'e'）-> exention（将 'n' 替换为 'x'）-> exection（将 'n' 替换为 'c'）-> execution（插入 'u'）。
```

**示例 3:**

```
输入: word1 = "", word2 = "abc"
输出: 3
解释: 需要插入三个字符将空字符串转换为 "abc"。
```

## 约束条件

- `0 <= word1.length, word2.length <= 500`
- `word1` 和 `word2` 由小写英文字母组成。

## 函数签名

```go
func minDistance(word1 string, word2 string) int
```
