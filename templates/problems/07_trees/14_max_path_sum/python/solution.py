from helpers import TreeNode, Optional


def solve_max_path_sum(root: Optional[TreeNode]) -> int:
    """Return the maximum path sum of any non-empty path in the tree.

    Time: O(n), Space: O(h).
    """
    result = float('-inf')

    def dfs(node: Optional[TreeNode]) -> int:
        nonlocal result
        if node is None:
            return 0
        left = dfs(node.left)
        if left < 0:
            left = 0
        right = dfs(node.right)
        if right < 0:
            right = 0
        path_sum = node.val + left + right
        if path_sum > result:
            result = path_sum
        return node.val + (left if left > right else right)

    dfs(root)
    return int(result)
