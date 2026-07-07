def solve_string_lines(input_data: str) -> str:
    """Read lines until EOF, reverse the word order in each non-empty line.

    Time: O(total input length), Space: O(total input length).
    """
    out_lines = []
    for line in input_data.splitlines():
        if not line.strip():
            continue
        words = line.split()
        out_lines.append(" ".join(reversed(words)))
    if not out_lines:
        return ""
    return "\n".join(out_lines) + "\n"
