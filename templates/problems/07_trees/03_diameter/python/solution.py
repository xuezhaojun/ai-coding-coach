from helpers import TreeNode, Optional


def solve_diameter_of_binary_tree(root: Optional[TreeNode]) -> int:
    """Return the length of the diameter of the tree.

    Time: O(n), Space: O(h).
    """
    result = 0

    def dfs(node: Optional[TreeNode]) -> int:
        nonlocal result
        if node is None:
            return 0
        left = dfs(node.left)
        right = dfs(node.right)
        if left + right > result:
            result = left + right
        return left + 1 if left > right else right + 1

    dfs(root)
    return result
