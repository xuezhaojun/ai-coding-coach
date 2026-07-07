from helpers import TreeNode, Optional


def solve_good_nodes(root: Optional[TreeNode]) -> int:
    """Return the number of good nodes in a binary tree.

    Time: O(n), Space: O(h).
    """
    def dfs(node: Optional[TreeNode], max_so_far: int) -> int:
        if node is None:
            return 0
        count = 0
        if node.val >= max_so_far:
            count = 1
            max_so_far = node.val
        return count + dfs(node.left, max_so_far) + dfs(node.right, max_so_far)

    return dfs(root, root.val)
