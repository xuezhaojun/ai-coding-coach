/**
 * Return the length of the longest increasing subsequence.
 * tails[i] is the smallest tail element for an increasing subsequence of length i+1.
 * Time: O(n log n), Space: O(n)
 */
export function solveLengthOfLIS(nums: number[]): number {
    const tails: number[] = [];
    for (const num of nums) {
        let lo = 0;
        let hi = tails.length;
        while (lo < hi) {
            const mid = (lo + hi) >> 1;
            if (tails[mid] < num) {
                lo = mid + 1;
            } else {
                hi = mid;
            }
        }
        if (lo === tails.length) {
            tails.push(num);
        } else {
            tails[lo] = num;
        }
    }
    return tails.length;
}
