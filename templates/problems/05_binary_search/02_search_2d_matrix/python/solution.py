def solve_search_matrix(matrix: list[list[int]], target: int) -> bool:
    """Treat the 2D matrix as a flat sorted array.

    Time: O(log(m*n)), Space: O(1).
    """
    if not matrix or not matrix[0]:
        return False
    m, n = len(matrix), len(matrix[0])
    lo, hi = 0, m * n - 1
    while lo <= hi:
        mid = lo + (hi - lo) // 2
        val = matrix[mid // n][mid % n]
        if val == target:
            return True
        elif val < target:
            lo = mid + 1
        else:
            hi = mid - 1
    return False
