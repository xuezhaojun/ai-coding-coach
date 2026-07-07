from helpers import TreeNode, Optional


def solve_max_depth(root: Optional[TreeNode]) -> int:
    """Return the maximum depth of a binary tree.

    Time: O(n), Space: O(h) where h is the height of the tree.
    """
    if root is None:
        return 0
    left = solve_max_depth(root.left)
    right = solve_max_depth(root.right)
    return left + 1 if left > right else right + 1
