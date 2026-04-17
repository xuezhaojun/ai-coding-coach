# 基于时间的键值存储

- **难度**：中等
- **分类**：二分查找
- **考点**：二分查找、哈希表、设计
- **链接**：[NeetCode](https://neetcode.io/problems/time-based-key-value-store) | [力扣 981](https://leetcode.cn/problems/time-based-key-value-store/)

## 题目描述

设计一个基于时间的键值数据结构，该结构可以在不同时间戳存储同一个键的多个值，并在特定时间戳检索键对应的值。

实现 `TimeMap` 类：
- `TimeMap()` 初始化数据结构对象。
- `Set(key string, value string, timestamp int)` 存储给定时间戳 `timestamp` 时的键 `key` 和值 `value`。
- `Get(key string, timestamp int) string` 返回一个值，该值满足之前调用了 `Set` 且 `timestamp_prev <= timestamp`。如果有多个这样的值，则返回 `timestamp_prev` 最大的那个值。如果没有满足条件的值，则返回空字符串 `""`。

对于给定的键，`Set` 的所有时间戳严格递增。

## 示例

**示例 1：**

```
输入：
["TimeMap", "Set", "Get", "Get", "Set", "Get", "Get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
输出：
[null, null, "bar", "bar", null, "bar2", "bar2"]
解释：
timeMap.Set("foo", "bar", 1);
timeMap.Get("foo", 1);  // 返回 "bar"
timeMap.Get("foo", 3);  // 返回 "bar"（最大的 timestamp <= 3 是 1）
timeMap.Set("foo", "bar2", 4);
timeMap.Get("foo", 4);  // 返回 "bar2"
timeMap.Get("foo", 5);  // 返回 "bar2"（最大的 timestamp <= 5 是 4）
```

**示例 2：**

```
输入：
["TimeMap", "Set", "Get"]
[[], ["key", "val", 5], ["key", 3]]
输出：
[null, null, ""]
解释：没有 timestamp <= 3 的值，所以返回 ""。
```

## 约束条件

- `1 <= key.length, value.length <= 100`
- `key` 和 `value` 由小写英文字母和数字组成
- `1 <= timestamp <= 10^7`
- 对于给定的键，`Set` 的所有时间戳严格递增
- 最多调用 `Set` 和 `Get` 操作 `2 * 10^5` 次

## 函数签名

```go
type TimeMap struct{}

func NewTimeMap() TimeMap
func (t *TimeMap) Set(key string, value string, timestamp int)
func (t *TimeMap) Get(key string, timestamp int) string
```
