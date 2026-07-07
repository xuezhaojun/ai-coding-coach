from helpers import TreeNode, Optional


def solve_invert_tree(root: Optional[TreeNode]) -> Optional[TreeNode]:
    """Invert a binary tree and return its root.

    Time: O(n), Space: O(h) where h is the height of the tree.
    """
    if root is None:
        return None
    root.left, root.right = root.right, root.left
    solve_invert_tree(root.left)
    solve_invert_tree(root.right)
    return root
