from collections import deque


def solve_alien_order(words: list[str]) -> str:
    """Determine character ordering from sorted alien words using topological sort.

    Time: O(C) where C is total characters across all words,
    Space: O(1) since alphabet is bounded.
    """
    # Build graph of character ordering
    graph: dict[str, set[str]] = {}
    in_degree: dict[str, int] = {}

    # Initialize all characters
    for w in words:
        for ch in w:
            if ch not in graph:
                graph[ch] = set()
                in_degree[ch] = 0

    # Compare adjacent words to find ordering constraints
    for i in range(len(words) - 1):
        w1, w2 = words[i], words[i + 1]
        # Check prefix violation
        if len(w1) > len(w2):
            prefix = True
            for j in range(len(w2)):
                if w1[j] != w2[j]:
                    prefix = False
                    break
            if prefix:
                return ""
        for j in range(min(len(w1), len(w2))):
            if w1[j] != w2[j]:
                if w2[j] not in graph[w1[j]]:
                    graph[w1[j]].add(w2[j])
                    in_degree[w2[j]] += 1
                break

    # BFS topological sort
    queue: deque[str] = deque(
        ch for ch in graph if in_degree[ch] == 0
    )

    result: list[str] = []
    while queue:
        ch = queue.popleft()
        result.append(ch)
        for nxt in graph[ch]:
            in_degree[nxt] -= 1
            if in_degree[nxt] == 0:
                queue.append(nxt)

    if len(result) != len(graph):
        return ""
    return "".join(result)
