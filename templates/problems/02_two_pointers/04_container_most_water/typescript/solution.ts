/**
 * Find the maximum area between two lines using two pointers.
 * Time: O(n), Space: O(1)
 */
export function solveMaxArea(height: number[]): number {
    let l = 0;
    let r = height.length - 1;
    let best = 0;
    while (l < r) {
        const h = Math.min(height[l], height[r]);
        const area = h * (r - l);
        if (area > best) {
            best = area;
        }
        if (height[l] < height[r]) {
            l++;
        } else {
            r--;
        }
    }
    return best;
}
