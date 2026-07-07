def solve_multiply(num1: str, num2: str) -> str:
    """Multiply two numbers represented as strings. O(m*n) time, O(m+n) space."""
    m, n = len(num1), len(num2)
    pos = [0] * (m + n)
    for i in range(m - 1, -1, -1):
        for j in range(n - 1, -1, -1):
            mul = (ord(num1[i]) - ord('0')) * (ord(num2[j]) - ord('0'))
            p1, p2 = i + j, i + j + 1
            total = mul + pos[p2]
            pos[p2] = total % 10
            pos[p1] += total // 10
    result: list[str] = []
    for v in pos:
        if len(result) == 0 and v == 0:
            continue
        result.append(str(v))
    if len(result) == 0:
        return "0"
    return "".join(result)
