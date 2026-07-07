def solve_sentinel_terminate(input_data: str) -> str:
    """Read pairs (a, b) until both are 0, return a+b for each pair (excluding sentinel).

    Time: O(n), Space: O(n) for output.
    """
    tokens = input_data.split()
    out_lines = []
    i = 0
    while i + 1 < len(tokens):
        a = int(tokens[i])
        b = int(tokens[i + 1])
        if a == 0 and b == 0:
            break
        out_lines.append(str(a + b))
        i += 2
    if not out_lines:
        return ""
    return "\n".join(out_lines) + "\n"
