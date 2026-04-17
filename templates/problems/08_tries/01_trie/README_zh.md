# 实现 Trie（前缀树）

- **难度**: 中等
- **分类**: 前缀树
- **考点**: 前缀树, 哈希表, 字符串
- **链接**: [NeetCode](https://neetcode.io/problems/implement-prefix-tree) | [力扣 208](https://leetcode.cn/problems/implement-trie-prefix-tree/)

## 题目描述

Trie（发音类似 "try"）或者说前缀树，是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。请实现 `Trie` 结构体，支持以下操作：

- `NewTrie()` 初始化前缀树对象。
- `Insert(word string)` 向前缀树中插入字符串 `word`。
- `Search(word string) bool` 如果字符串 `word` 在前缀树中（即之前已被插入），返回 `true`；否则返回 `false`。
- `StartsWith(prefix string) bool` 如果之前已插入的字符串中存在以 `prefix` 为前缀的字符串，返回 `true`；否则返回 `false`。

## 示例

**示例 1:**

```
输入:
  NewTrie()
  Insert("apple")
  Search("apple")   -> true
  Search("app")     -> false
  StartsWith("app") -> true
  Insert("app")
  Search("app")     -> true

解释: 插入 "apple" 后，搜索 "apple" 返回 true，但 "app" 返回 false，因为它没有作为完整单词被插入。然而 StartsWith("app") 返回 true，因为 "apple" 有这个前缀。插入 "app" 后，搜索 "app" 返回 true。
```

**示例 2:**

```
输入:
  NewTrie()
  Insert("the")
  Insert("then")
  Insert("them")
  Search("the")      -> true
  Search("they")     -> false
  StartsWith("the")  -> true
  StartsWith("thx")  -> false

解释: 单词 "the"、"then" 和 "them" 共享前缀 "the"。搜索 "they" 失败，因为从未插入过。
```

## 约束条件

- `1 <= word.length, prefix.length <= 2000`
- `word` 和 `prefix` 仅由小写英文字母组成。
- `Insert`、`Search` 和 `StartsWith` 的调用次数总计不超过 `3 * 10^4` 次。

## 函数签名

```go
type Trie struct {
    Root *TrieNode
}

func NewTrie() Trie

func (t *Trie) Insert(word string)

func (t *Trie) Search(word string) bool

func (t *Trie) StartsWith(prefix string) bool
```
