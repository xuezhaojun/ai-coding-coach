def solve_letter_combinations(digits: str) -> list[str]:
    """Generate all letter combinations for phone digits.

    Time: O(4^n * n), Space: O(n) recursion depth
    """
    if len(digits) == 0:
        return []
    phone = {
        "2": "abc", "3": "def", "4": "ghi", "5": "jkl",
        "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz",
    }
    result: list[str] = []

    def backtrack(idx: int, current: list[str]) -> None:
        if idx == len(digits):
            result.append("".join(current))
            return
        for ch in phone[digits[idx]]:
            current.append(ch)
            backtrack(idx + 1, current)
            current.pop()

    backtrack(0, [])
    return result
