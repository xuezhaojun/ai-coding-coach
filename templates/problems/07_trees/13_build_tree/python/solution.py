from helpers import TreeNode, Optional


def solve_build_tree_from_pre_in(preorder: list[int], inorder: list[int]) -> Optional[TreeNode]:
    """Construct a binary tree from preorder and inorder traversals.

    Time: O(n), Space: O(n).
    """
    in_map = {v: i for i, v in enumerate(inorder)}
    idx = 0

    def build(lo: int, hi: int) -> Optional[TreeNode]:
        nonlocal idx
        if lo > hi:
            return None
        root_val = preorder[idx]
        idx += 1
        node = TreeNode(root_val)
        mid = in_map[root_val]
        node.left = build(lo, mid - 1)
        node.right = build(mid + 1, hi)
        return node

    return build(0, len(inorder) - 1)
