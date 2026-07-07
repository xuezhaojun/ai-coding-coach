from collections import deque


def solve_ladder_length(begin_word: str, end_word: str, word_list: list[str]) -> int:
    """Find shortest transformation sequence using BFS.

    Time: O(n * m * 26) where n=wordList length, m=word length,
    Space: O(n * m).
    """
    word_set = set(word_list)
    if end_word not in word_set:
        return 0
    if begin_word == end_word:
        return 0

    queue: deque[str] = deque([begin_word])
    visited: set[str] = {begin_word}
    length = 1

    while queue:
        size = len(queue)
        for _ in range(size):
            word = queue.popleft()
            chars = list(word)
            for j in range(len(chars)):
                original = chars[j]
                for c in range(ord("a"), ord("z") + 1):
                    ch = chr(c)
                    if ch == original:
                        continue
                    chars[j] = ch
                    nxt = "".join(chars)
                    if nxt == end_word:
                        return length + 1
                    if nxt in word_set and nxt not in visited:
                        visited.add(nxt)
                        queue.append(nxt)
                chars[j] = original
        length += 1
    return 0
