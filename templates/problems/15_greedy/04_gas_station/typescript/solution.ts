/**
 * Find the starting gas station for a circular route. O(n) time, O(1) space.
 */
export function solveCanCompleteCircuit(gas: number[], cost: number[]): number {
    let totalSurplus = 0;
    let currentSurplus = 0;
    let start = 0;
    for (let i = 0; i < gas.length; i++) {
        totalSurplus += gas[i] - cost[i];
        currentSurplus += gas[i] - cost[i];
        if (currentSurplus < 0) {
            start = i + 1;
            currentSurplus = 0;
        }
    }
    if (totalSurplus < 0) return -1;
    return start;
}
