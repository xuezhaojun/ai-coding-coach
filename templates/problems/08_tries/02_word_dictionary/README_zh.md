# 添加与搜索单词 - 数据结构设计

- **难度**: 中等
- **分类**: 前缀树
- **考点**: 前缀树, 深度优先搜索, 回溯, 通配符匹配
- **链接**: [NeetCode](https://neetcode.io/problems/design-word-search-data-structure) | [力扣 211](https://leetcode.cn/problems/design-add-and-search-words-data-structure/)

## 题目描述

设计一个数据结构，支持添加新单词和查找字符串是否与任何先前添加的字符串匹配。该数据结构需要支持以下操作：

- `AddWord(word string)` 将 `word` 添加到数据结构中，以便后续匹配。
- `Search(word string) bool` 如果数据结构中存在与 `word` 匹配的字符串，返回 `true`。`word` 中可能包含点号 `'.'`，每个点号可以匹配任意一个字母。

## 示例

**示例 1:**

```
输入:
  NewWordDictionary()
  AddWord("bad")
  AddWord("dad")
  AddWord("mad")
  Search("pad")  -> false
  Search("bad")  -> true
  Search(".ad")  -> true
  Search("b..")  -> true

解释: "pad" 从未被添加。"bad" 完全匹配。".ad" 可以匹配 "bad"、"dad" 或 "mad"，因为 '.' 匹配任意字符。"b.." 匹配 "bad"。
```

**示例 2:**

```
输入:
  NewWordDictionary()
  AddWord("a")
  Search(".")   -> true
  Search("a")   -> true
  Search("..")  -> false

解释: "." 匹配 "a"（单个字符）。".." 需要一个长度为 2 的单词，但只添加了 "a"（长度为 1）。
```

## 约束条件

- `1 <= word.length <= 25`
- `AddWord` 中的 `word` 仅由小写英文字母组成。
- `Search` 中的 `word` 由小写英文字母或 `'.'` 组成。
- `AddWord` 和 `Search` 的调用次数总计不超过 `10^4` 次。

## 函数签名

```go
type WordDictionary struct {
    Root *TrieNode
}

func NewWordDictionary() WordDictionary

func (wd *WordDictionary) AddWord(word string)

func (wd *WordDictionary) Search(word string) bool
```
