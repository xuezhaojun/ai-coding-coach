def solve_product_except_self(nums: list[int]) -> list[int]:
    """Compute product array without division using prefix/suffix passes.

    Time: O(n), Space: O(1) extra (output array not counted)
    """
    n = len(nums)
    result = [0] * n

    # Forward pass: result[i] = product of all elements before i
    result[0] = 1
    for i in range(1, n):
        result[i] = result[i - 1] * nums[i - 1]

    # Backward pass: multiply by product of all elements after i
    suffix = 1
    for i in range(n - 2, -1, -1):
        suffix *= nums[i + 1]
        result[i] *= suffix

    return result
