# 最小栈

- **难度**：中等
- **分类**：栈
- **考点**：栈、设计
- **链接**：[NeetCode](https://neetcode.io/problems/minimum-stack) | [力扣 155](https://leetcode.cn/problems/min-stack/)

## 题目描述

设计一个支持 `push`、`pop`、`top` 操作，并能在常数时间内检索到最小元素的栈。

实现 `MinStack` 类：
- `MinStack()` 初始化堆栈对象。
- `Push(val int)` 将元素 `val` 推入堆栈。
- `Pop()` 删除堆栈顶部的元素。
- `Top() int` 获取堆栈顶部的元素。
- `GetMin() int` 获取堆栈中的最小元素。

每个函数的时间复杂度必须为 O(1)。

## 示例

**示例 1：**

```
输入：
["MinStack","Push","Push","Push","GetMin","Pop","Top","GetMin"]
[[],[-2],[0],[-3],[],[],[],[]]
输出：
[null,null,null,null,-3,null,0,-2]
解释：
MinStack minStack = new MinStack();
minStack.Push(-2);
minStack.Push(0);
minStack.Push(-3);
minStack.GetMin(); // 返回 -3
minStack.Pop();
minStack.Top();    // 返回 0
minStack.GetMin(); // 返回 -2
```

## 约束条件

- `-2^31 <= val <= 2^31 - 1`
- `Pop`、`Top` 和 `GetMin` 操作总是在非空栈上调用
- 最多调用 `3 * 10^4` 次 `Push`、`Pop`、`Top` 和 `GetMin`

## 函数签名

```go
type MinStack struct{}

func NewMinStack() MinStack
func (s *MinStack) Push(val int)
func (s *MinStack) Pop()
func (s *MinStack) Top() int
func (s *MinStack) GetMin() int
```
