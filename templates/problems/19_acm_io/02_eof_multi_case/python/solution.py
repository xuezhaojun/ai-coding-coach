def solve_eof_multi_case(input_data: str) -> str:
    """Read pairs (a, b) until EOF, return a+b for each pair.

    Time: O(n), Space: O(n) for output.
    """
    tokens = input_data.split()
    out_lines = []
    i = 0
    while i + 1 < len(tokens):
        a = int(tokens[i])
        b = int(tokens[i + 1])
        out_lines.append(str(a + b))
        i += 2
    if not out_lines:
        return ""
    return "\n".join(out_lines) + "\n"
