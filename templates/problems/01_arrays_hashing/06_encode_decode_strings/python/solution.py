def solve_encode(strs: list[str]) -> str:
    """Encode a list of strings to a single string using length-prefixed format.

    Time: O(n), Space: O(n) where n is total characters across all strings
    """
    parts = []
    for s in strs:
        parts.append(str(len(s)))
        parts.append("#")
        parts.append(s)
    return "".join(parts)


def solve_decode(s: str) -> list[str]:
    """Decode a single string back to a list of strings.

    Time: O(n), Space: O(n)
    """
    result: list[str] = []
    i = 0
    while i < len(s):
        j = i
        while s[j] != "#":
            j += 1
        length = int(s[i:j])
        result.append(s[j + 1:j + 1 + length])
        i = j + 1 + length
    return result
