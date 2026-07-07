/**
 * Merge a new interval into a sorted list. O(n) time, O(n) space.
 */
export function solveInsert(intervals: number[][], newInterval: number[]): number[][] {
    const result: number[][] = [];
    let i = 0;
    // Add all intervals before the new interval
    while (i < intervals.length && intervals[i][1] < newInterval[0]) {
        result.push(intervals[i]);
        i++;
    }
    // Merge overlapping intervals with newInterval
    while (i < intervals.length && intervals[i][0] <= newInterval[1]) {
        if (intervals[i][0] < newInterval[0]) {
            newInterval[0] = intervals[i][0];
        }
        if (intervals[i][1] > newInterval[1]) {
            newInterval[1] = intervals[i][1];
        }
        i++;
    }
    result.push(newInterval);
    // Add remaining intervals
    while (i < intervals.length) {
        result.push(intervals[i]);
        i++;
    }
    return result;
}
