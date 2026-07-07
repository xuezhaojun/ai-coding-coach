from helpers import TreeNode, Optional


def solve_kth_smallest(root: Optional[TreeNode], k: int) -> int:
    """Return the kth smallest value in a BST.

    Time: O(h + k), Space: O(h).
    """
    count = 0
    result = 0

    def inorder(node: Optional[TreeNode]) -> None:
        nonlocal count, result
        if node is None or count >= k:
            return
        inorder(node.left)
        count += 1
        if count == k:
            result = node.val
            return
        inorder(node.right)

    inorder(root)
    return result
