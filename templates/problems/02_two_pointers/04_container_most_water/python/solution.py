def solve_max_area(height: list[int]) -> int:
    """Find the maximum area between two lines using two pointers.

    Time: O(n), Space: O(1)
    """
    l, r = 0, len(height) - 1
    best = 0
    while l < r:
        h = min(height[l], height[r])
        area = h * (r - l)
        if area > best:
            best = area
        if height[l] < height[r]:
            l += 1
        else:
            r -= 1
    return best
