def solve_check_valid_string(s: str) -> bool:
    """Track min/max open count range. O(n) time, O(1) space."""
    lo, hi = 0, 0  # range of possible open paren counts
    for c in s:
        if c == "(":
            lo += 1
            hi += 1
        elif c == ")":
            lo -= 1
            hi -= 1
        else:  # '*'
            lo -= 1  # * treated as )
            hi += 1  # * treated as (
        if hi < 0:
            return False
        if lo < 0:
            lo = 0
    return lo == 0
