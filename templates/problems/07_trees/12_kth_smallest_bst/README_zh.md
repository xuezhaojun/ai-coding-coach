# 二叉搜索树中第 K 小的元素

- **难度**: 中等
- **分类**: 树
- **考点**: 二叉搜索树, 中序遍历, 深度优先搜索
- **链接**: [NeetCode](https://neetcode.io/problems/kth-smallest-integer-in-bst) | [力扣 230](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)

## 题目描述

给定一个二叉搜索树的根节点和一个整数 `k`，返回树中所有节点值中第 `k` 小的值（1-indexed）。由于 BST 性质保证中序遍历按升序访问节点，因此中序遍历中第 `k` 个访问的元素即为答案。

## 示例

**示例 1:**

```
输入: root = [3,1,4,null,2], k = 1
输出: 1
解释: 中序遍历为 [1, 2, 3, 4]。第 1 小的元素是 1。
```

**示例 2:**

```
输入: root = [5,3,6,2,4,null,null,1], k = 3
输出: 3
解释: 中序遍历为 [1, 2, 3, 4, 5, 6]。第 3 小的元素是 3。
```

**示例 3:**

```
输入: root = [3,1,4,null,2], k = 2
输出: 2
解释: 中序遍历为 [1, 2, 3, 4]。第 2 小的元素是 2。
```

## 约束条件

- 树中节点数为 `n`。
- `1 <= k <= n <= 10^4`
- `0 <= Node.val <= 10^4`

## 函数签名

```go
func kthSmallest(root *TreeNode, k int) int
```
