def solve_graph_input(input_data: str) -> str:
    """Read a directed graph with n nodes and m edges, return the out-degree
    of each node (1..n), one per line.

    Time: O(n + m), Space: O(n).
    """
    tokens = input_data.split()
    idx = 0
    n = int(tokens[idx]); idx += 1
    m = int(tokens[idx]); idx += 1

    degree = [0] * (n + 1)
    for _ in range(m):
        u = int(tokens[idx]); idx += 1
        v = int(tokens[idx]); idx += 1  # noqa: F841 (v not needed for out-degree)
        degree[u] += 1

    out_lines = [str(degree[i]) for i in range(1, n + 1)]
    return "\n".join(out_lines) + "\n"
