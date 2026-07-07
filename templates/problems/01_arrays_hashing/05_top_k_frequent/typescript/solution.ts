/**
 * Find k most frequent elements using bucket sort.
 * Time: O(n), Space: O(n)
 */
export function solveTopKFrequent(nums: number[], k: number): number[] {
    const freq = new Map<number, number>();
    for (const n of nums) {
        freq.set(n, (freq.get(n) ?? 0) + 1);
    }

    // Bucket sort: index = frequency, value = list of numbers with that frequency
    const buckets: number[][] = Array.from({ length: nums.length + 1 }, () => []);
    for (const [num, count] of freq) {
        buckets[count].push(num);
    }

    const result: number[] = [];
    for (let i = buckets.length - 1; i >= 0 && result.length < k; i--) {
        result.push(...buckets[i]);
    }
    return result.slice(0, k);
}
