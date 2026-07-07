/**
 * Check if target triplet can be formed. O(n) time, O(1) space.
 */
export function solveMergeTriplets(triplets: number[][], target: number[]): boolean {
    const good = [false, false, false];
    for (const t of triplets) {
        if (t[0] > target[0] || t[1] > target[1] || t[2] > target[2]) continue;
        for (let i = 0; i < 3; i++) {
            if (t[i] === target[i]) good[i] = true;
        }
    }
    return good[0] && good[1] && good[2];
}
