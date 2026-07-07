/**
 * Compute product array without division using prefix/suffix passes.
 * Time: O(n), Space: O(1) extra (output array not counted)
 */
export function solveProductExceptSelf(nums: number[]): number[] {
    const n = nums.length;
    const result = new Array(n);

    // Forward pass: result[i] = product of all elements before i
    result[0] = 1;
    for (let i = 1; i < n; i++) {
        result[i] = result[i - 1] * nums[i - 1];
    }

    // Backward pass: multiply by product of all elements after i
    let suffix = 1;
    for (let i = n - 2; i >= 0; i--) {
        suffix *= nums[i + 1];
        result[i] *= suffix;
    }

    return result;
}
