from helpers import TreeNode, Optional


def solve_serialize(root: Optional[TreeNode]) -> str:
    """Encode a binary tree to a single string.

    Time: O(n), Space: O(n).
    """
    parts: list[str] = []

    def preorder(node: Optional[TreeNode]) -> None:
        if node is None:
            parts.append("N")
            return
        parts.append(str(node.val))
        preorder(node.left)
        preorder(node.right)

    preorder(root)
    return ",".join(parts)


def solve_deserialize(data: str) -> Optional[TreeNode]:
    """Decode the encoded string to a binary tree.

    Time: O(n), Space: O(n).
    """
    tokens = data.split(",")
    idx = 0

    def build() -> Optional[TreeNode]:
        nonlocal idx
        if idx >= len(tokens) or tokens[idx] == "N" or tokens[idx] == "":
            idx += 1
            return None
        val = int(tokens[idx])
        idx += 1
        node = TreeNode(val)
        node.left = build()
        node.right = build()
        return node

    return build()
