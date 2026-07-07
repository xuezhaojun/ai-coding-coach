from __future__ import annotations


class _SolveWDNode:
    __slots__ = ("children", "is_end")

    def __init__(self):
        self.children: dict[str, _SolveWDNode] = {}
        self.is_end = False


class SolveWordDictionary:
    """Supports adding words and searching with '.' wildcards."""

    def __init__(self):
        self.root = _SolveWDNode()

    def add_word(self, word: str) -> None:
        node = self.root
        for c in word:
            if c not in node.children:
                node.children[c] = _SolveWDNode()
            node = node.children[c]
        node.is_end = True

    def search(self, word: str) -> bool:
        return _search_node(self.root, word, 0)


def _search_node(node: _SolveWDNode, word: str, i: int) -> bool:
    if i == len(word):
        return node.is_end
    c = word[i]
    if c == ".":
        for child in node.children.values():
            if _search_node(child, word, i + 1):
                return True
        return False
    child = node.children.get(c)
    if child is None:
        return False
    return _search_node(child, word, i + 1)
