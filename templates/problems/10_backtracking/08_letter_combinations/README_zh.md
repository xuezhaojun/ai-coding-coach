# 电话号码的字母组合

- **难度**: 中等
- **分类**: 回溯
- **考点**: 回溯, 字符串, 哈希表
- **链接**: [NeetCode](https://neetcode.io/problems/letter-combinations-of-a-phone-number) | [力扣 17](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)

## 题目描述

给定一个仅包含数字 `2-9` 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。

数字到字母的映射与电话按键相同：2 对应 "abc"，3 对应 "def"，4 对应 "ghi"，5 对应 "jkl"，6 对应 "mno"，7 对应 "pqrs"，8 对应 "tuv"，9 对应 "wxyz"。注意数字 1 不对应任何字母。

## 示例

**示例 1:**

```
输入: digits = "23"
输出: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**示例 2:**

```
输入: digits = ""
输出: []
```

**示例 3:**

```
输入: digits = "2"
输出: ["a","b","c"]
```

## 约束条件

- `0 <= digits.length <= 4`
- `digits[i]` 是范围 `['2', '9']` 的一个数字。

## 函数签名

```go
func letterCombinations(digits string) []string
```
