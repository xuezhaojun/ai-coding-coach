from __future__ import annotations


class _TrieNode:
    __slots__ = ("children", "word")

    def __init__(self):
        self.children: list[_TrieNode | None] = [None] * 26
        self.word: str = ""


def solve_find_words(board: list[list[str]], words: list[str]) -> list[str]:
    """Return all words formed by sequentially adjacent cells on the board.

    Time: O(m*n * 4^L) where L is max word length, Space: O(total chars in words).
    """
    root = _TrieNode()
    for w in words:
        node = root
        for ch in w:
            idx = ord(ch) - ord("a")
            if node.children[idx] is None:
                node.children[idx] = _TrieNode()
            node = node.children[idx]
        node.word = w

    rows = len(board)
    cols = len(board[0]) if rows > 0 else 0
    result: list[str] = []

    def dfs(r: int, c: int, node: _TrieNode) -> None:
        if r < 0 or r >= rows or c < 0 or c >= cols or board[r][c] == "#":
            return
        ch = board[r][c]
        nxt = node.children[ord(ch) - ord("a")]
        if nxt is None:
            return
        if nxt.word != "":
            result.append(nxt.word)
            nxt.word = ""
        board[r][c] = "#"
        dfs(r + 1, c, nxt)
        dfs(r - 1, c, nxt)
        dfs(r, c + 1, nxt)
        dfs(r, c - 1, nxt)
        board[r][c] = ch

    for r in range(rows):
        for c in range(cols):
            dfs(r, c, root)
    return result
