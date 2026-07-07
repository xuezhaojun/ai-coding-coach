def solve_largest_rectangle_area(heights: list[int]) -> int:
    """Use a monotonic increasing stack.

    Time: O(n), Space: O(n)
    """
    stack: list[int] = []  # stack of indices
    max_area = 0
    for i in range(len(heights) + 1):
        h = heights[i] if i < len(heights) else 0
        while stack and h < heights[stack[-1]]:
            height = heights[stack.pop()]
            width = i if not stack else i - stack[-1] - 1
            area = height * width
            if area > max_area:
                max_area = area
        stack.append(i)
    return max_area
