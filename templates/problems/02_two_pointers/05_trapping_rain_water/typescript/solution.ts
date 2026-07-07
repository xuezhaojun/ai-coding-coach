/**
 * Compute trapped rainwater using two pointers.
 * Time: O(n), Space: O(1)
 */
export function solveTrap(height: number[]): number {
    if (height.length === 0) {
        return 0;
    }
    let l = 0;
    let r = height.length - 1;
    let leftMax = height[l];
    let rightMax = height[r];
    let water = 0;

    while (l < r) {
        if (leftMax < rightMax) {
            l++;
            if (height[l] > leftMax) {
                leftMax = height[l];
            } else {
                water += leftMax - height[l];
            }
        } else {
            r--;
            if (height[r] > rightMax) {
                rightMax = height[r];
            } else {
                water += rightMax - height[r];
            }
        }
    }
    return water;
}
