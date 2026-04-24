# 二叉树的序列化与反序列化

- **难度**: 困难
- **分类**: 树
- **考点**: 二叉树, 前序遍历, 字符串处理
- **链接**: [NeetCode](https://neetcode.io/problems/serialize-and-deserialize-binary-tree) | [力扣 297](https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/)

## 题目描述

设计一个算法来序列化和反序列化二叉树。序列化是将树转换为字符串表示以便存储或传输的过程，反序列化是从字符串重建树的过程。对序列化/反序列化算法的实现方式没有限制，只要二叉树能被序列化为字符串，并且该字符串能被反序列化回原始树结构即可。

![serialize-and-deserialize-binary-tree](assets/serdeser.jpg)

## 示例

**示例 1:**

```
输入: root = [1,2,3,null,null,4,5]
输出: [1,2,3,null,null,4,5]
解释: 树被序列化为字符串，然后反序列化回来产生一棵完全相同的树。
```

**示例 2:**

```
输入: root = []
输出: []
解释: 空树的序列化和反序列化结果为空。
```

**示例 3:**

```
输入: root = [1]
输出: [1]
解释: 单节点树经过序列化和反序列化后被保留。
```

## 约束条件

- 树中节点数目范围在 `[0, 10^4]` 内。
- `-1000 <= Node.val <= 1000`

## 函数签名

```go
func serialize(root *TreeNode) string

func deserialize(data string) *TreeNode
```
