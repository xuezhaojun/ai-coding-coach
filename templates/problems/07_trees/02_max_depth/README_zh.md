# 二叉树的最大深度

- **难度**: 简单
- **分类**: 树
- **考点**: 二叉树, 递归, 深度优先搜索
- **链接**: [NeetCode](https://neetcode.io/problems/depth-of-binary-tree) | [力扣 104](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)

## 题目描述

给定一棵二叉树的根节点，返回其最大深度。二叉树的最大深度是从根节点到最远叶子节点的最长路径上的节点数。叶子节点是没有子节点的节点。

![maximum-depth-of-binary-tree](assets/tmp-tree.jpg)

## 示例

**示例 1:**

```
输入: root = [3,9,20,null,null,15,7]
输出: 3
解释: 最长路径为 根节点(3) -> 右子节点(20) -> 左子节点(15) 或 右子节点(7)，共 3 个节点。
```

**示例 2:**

```
输入: root = [1,null,2]
输出: 2
解释: 树只有一个右子节点，因此深度为 2。
```

**示例 3:**

```
输入: root = []
输出: 0
解释: 空树的深度为 0。
```

## 约束条件

- 树中节点数目范围在 `[0, 10^4]` 内。
- `-100 <= Node.val <= 100`

## 函数签名

```go
func maxDepth(root *TreeNode) int
```
