# 单词拆分

- **难度**: 中等
- **分类**: 一维动态规划
- **考点**: 动态规划, 字符串, 哈希表
- **链接**: [NeetCode](https://neetcode.io/problems/word-break) | [力扣 139](https://leetcode.cn/problems/word-break/)

## 题目描述

给你一个字符串 `s` 和一个字符串列表 `wordDict` 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 `s`，则返回 `true`。注意，字典中的同一个单词可以在分割中被重复使用。

## 示例

**示例 1:**

```
输入: s = "leetcode", wordDict = ["leet","code"]
输出: true
解释: "leetcode" 可以被拆分为 "leet code"。
```

**示例 2:**

```
输入: s = "applepenapple", wordDict = ["apple","pen"]
输出: true
解释: "applepenapple" 可以被拆分为 "apple pen apple"。注意单词 "apple" 被重复使用。
```

**示例 3:**

```
输入: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
输出: false
```

## 约束条件

- `1 <= s.length <= 300`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 20`
- `s` 和 `wordDict[i]` 仅由小写英文字母组成。
- `wordDict` 中的所有字符串互不相同。

## 函数签名

```go
func wordBreak(s string, wordDict []string) bool
```
