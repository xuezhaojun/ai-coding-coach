# 数据流中的第 K 大元素

- **难度**: 简单
- **分类**: 堆
- **考点**: 最小堆, 优先队列, 流式处理
- **链接**: [NeetCode](https://neetcode.io/problems/kth-largest-integer-in-a-stream) | [力扣 703](https://leetcode.cn/problems/kth-largest-element-in-a-stream/)

## 题目描述

设计一个找到数据流中第 `k` 大元素的类。注意是排序后的第 `k` 大元素，而非第 `k` 个不同的元素。该类使用整数 `k` 和整数数组 `nums` 初始化。每次调用 `Add` 方法时，向流中插入一个新值，并返回当前流中第 `k` 大的元素。

## 示例

**示例 1:**

```
输入:
  Constructor(3, [4, 5, 8, 2])
  Add(3)  -> 4
  Add(5)  -> 5
  Add(10) -> 5
  Add(9)  -> 8
  Add(4)  -> 8

解释: 初始化后，排序后的流为 [2,4,5,8]。第 3 大元素是 4。
  Add(3): 流为 [2,3,4,5,8]，第 3 大 = 4。
  Add(5): 流为 [2,3,4,5,5,8]，第 3 大 = 5。
  Add(10): 流为 [2,3,4,5,5,8,10]，第 3 大 = 5。
  Add(9): 流为 [2,3,4,5,5,8,9,10]，第 3 大 = 8。
  Add(4): 流为 [2,3,4,4,5,5,8,9,10]，第 3 大 = 8。
```

**示例 2:**

```
输入:
  Constructor(1, [])
  Add(-1) -> -1
  Add(1)  -> 1
  Add(-2) -> 1

解释: 当 k=1 时，第 1 大元素始终是流中的最大值。
```

## 约束条件

- `1 <= k <= 10^4`
- `0 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `-10^4 <= val <= 10^4`
- `Add` 最多被调用 `10^4` 次。
- 保证搜索第 `k` 大元素时流中至少有 `k` 个元素。

## 函数签名

```go
type KthLargest struct {}

func Constructor(k int, nums []int) KthLargest

func (kl *KthLargest) Add(val int) int
```
