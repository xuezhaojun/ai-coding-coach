# LRU 缓存

- **难度**：中等
- **分类**：链表
- **考点**：链表、哈希表、设计、双向链表
- **链接**：[NeetCode](https://neetcode.io/problems/lru-cache) | [力扣 146](https://leetcode.cn/problems/lru-cache/)

## 题目描述

请你设计并实现一个满足 LRU（最近最少使用）缓存约束的数据结构。

实现 `LRUCache` 类：
- `Constructor(capacity int) LRUCache` 以正整数作为容量 `capacity` 初始化 LRU 缓存。
- `Get(key int) int` 如果关键字 `key` 存在于缓存中，则返回关键字的值，否则返回 `-1`。
- `Put(key int, value int)` 如果关键字 `key` 已经存在，则变更其数据值 `value`；如果不存在，则向缓存中插入该组 `key-value`。如果插入操作导致关键字数量超过 `capacity`，则应该逐出最久未使用的关键字。

函数 `Get` 和 `Put` 必须以 O(1) 的平均时间复杂度运行。

## 示例

**示例 1：**

```
输入：
["LRUCache", "Put", "Put", "Get", "Put", "Get", "Put", "Get", "Get", "Get"]
[[2], [1,1], [2,2], [1], [3,3], [2], [4,4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, null, -1, 3, 4]
解释：
LRUCache lRUCache = new LRUCache(2);
lRUCache.Put(1, 1);  // 缓存是 {1=1}
lRUCache.Put(2, 2);  // 缓存是 {1=1, 2=2}
lRUCache.Get(1);      // 返回 1
lRUCache.Put(3, 3);  // 逐出键 2，缓存是 {1=1, 3=3}
lRUCache.Get(2);      // 返回 -1（未找到）
lRUCache.Put(4, 4);  // 逐出键 1，缓存是 {4=4, 3=3}
lRUCache.Get(1);      // 返回 -1（未找到）
lRUCache.Get(3);      // 返回 3
lRUCache.Get(4);      // 返回 4
```

## 约束条件

- `1 <= capacity <= 3000`
- `0 <= key <= 10^4`
- `0 <= value <= 10^5`
- 最多调用 `2 * 10^5` 次 `Get` 和 `Put`

## 函数签名

```go
type LRUCache struct{}

func Constructor(capacity int) LRUCache
func (c *LRUCache) Get(key int) int
func (c *LRUCache) Put(key int, value int)
```
