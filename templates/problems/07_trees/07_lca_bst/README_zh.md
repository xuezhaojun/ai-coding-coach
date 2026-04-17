# 二叉搜索树的最近公共祖先

- **难度**: 中等
- **分类**: 树
- **考点**: 二叉搜索树, BST 性质, 迭代
- **链接**: [NeetCode](https://neetcode.io/problems/lowest-common-ancestor-in-binary-search-tree) | [力扣 235](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/)

## 题目描述

给定一个二叉搜索树（BST）以及两个节点 `p` 和 `q`，找到它们的最近公共祖先（LCA）。最近公共祖先是指在树中同时拥有 `p` 和 `q` 作为后代的最深节点（一个节点也可以是它自己的后代）。BST 中所有节点的值唯一，且 `p` 和 `q` 保证存在于树中。

## 示例

**示例 1:**

```
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是根节点 6，因为 2 在左子树，8 在右子树。
```

**示例 2:**

```
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是节点 2，因为 4 是 2 的后代，且一个节点可以是自己的祖先。
```

**示例 3:**

```
输入: root = [2,1,3], p = 1, q = 3
输出: 2
解释: 节点 1 和节点 3 的最近公共祖先是根节点 2。
```

## 约束条件

- 树中节点数目范围在 `[2, 10^5]` 内。
- `-10^9 <= Node.val <= 10^9`
- 所有 `Node.val` 互不相同。
- `p != q`
- `p` 和 `q` 均存在于 BST 中。

## 函数签名

```go
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode
```
