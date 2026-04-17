# 数据流的中位数

- **难度**: 困难
- **分类**: 堆
- **考点**: 双堆, 最大堆, 最小堆, 设计
- **链接**: [NeetCode](https://neetcode.io/problems/find-median-in-a-data-stream) | [力扣 295](https://leetcode.cn/problems/find-median-from-data-stream/)

## 题目描述

中位数是有序整数列表中间的数。如果列表长度是偶数，没有中间值，则中位数是中间两个数的平均值。设计一个数据结构，支持以下两种操作：

- `AddNum(num int)` 从数据流中添加一个整数 `num` 到数据结构中。
- `FindMedian() float64` 返回目前所有已添加元素的中位数。

## 示例

**示例 1:**

```
输入:
  ConstructorMedianFinder()
  AddNum(1)
  AddNum(2)
  FindMedian() -> 1.5
  AddNum(3)
  FindMedian() -> 2.0

解释: 添加 1 和 2 后，有序列表为 [1,2]。中位数为 (1+2)/2 = 1.5。添加 3 后，有序列表为 [1,2,3]。中位数为 2.0。
```

**示例 2:**

```
输入:
  ConstructorMedianFinder()
  AddNum(5)
  FindMedian() -> 5.0

解释: 只有一个元素时，中位数就是该元素本身。
```

**示例 3:**

```
输入:
  ConstructorMedianFinder()
  AddNum(5)
  AddNum(4)
  AddNum(3)
  AddNum(2)
  AddNum(1)
  FindMedian() -> 3.0

解释: 添加所有元素后，有序列表为 [1,2,3,4,5]。中位数是中间值 3。
```

## 约束条件

- `-10^5 <= num <= 10^5`
- 调用 `FindMedian` 之前，数据结构中至少有一个元素。
- `AddNum` 和 `FindMedian` 的调用次数总计不超过 `5 * 10^4` 次。

## 函数签名

```go
type MedianFinder struct {}

func ConstructorMedianFinder() MedianFinder

func (mf *MedianFinder) AddNum(num int)

func (mf *MedianFinder) FindMedian() float64
```
