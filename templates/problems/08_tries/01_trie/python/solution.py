from __future__ import annotations


class _SolveTrieNode:
    __slots__ = ("children", "is_end")

    def __init__(self):
        self.children: dict[str, _SolveTrieNode] = {}
        self.is_end = False


class SolveTrie:
    """A prefix tree. Time per op: O(m) where m is word length."""

    def __init__(self):
        self.root = _SolveTrieNode()

    def insert(self, word: str) -> None:
        node = self.root
        for c in word:
            if c not in node.children:
                node.children[c] = _SolveTrieNode()
            node = node.children[c]
        node.is_end = True

    def search(self, word: str) -> bool:
        node = self._find_node(word)
        return node is not None and node.is_end

    def starts_with(self, prefix: str) -> bool:
        return self._find_node(prefix) is not None

    def _find_node(self, prefix: str) -> _SolveTrieNode | None:
        node = self.root
        for c in prefix:
            node = node.children.get(c)
            if node is None:
                return None
        return node
