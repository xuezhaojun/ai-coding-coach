def solve_fixed_format(input_data: str) -> str:
    """Read n and n integers, return their sum.

    Time: O(n), Space: O(n) for tokenizing input.
    """
    tokens = input_data.split()
    n = int(tokens[0])
    total = 0
    for i in range(1, n + 1):
        total += int(tokens[i])
    return f"{total}\n"
