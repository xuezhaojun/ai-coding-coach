/**
 * Use a sweep line approach. O(n log n) time, O(n) space.
 */
export function solveMinMeetingRooms(intervals: number[][]): number {
    if (intervals.length === 0) {
        return 0;
    }
    const starts = intervals.map(iv => iv[0]).sort((a, b) => a - b);
    const ends = intervals.map(iv => iv[1]).sort((a, b) => a - b);

    let rooms = 0, maxRooms = 0;
    let s = 0, e = 0;
    while (s < starts.length) {
        if (starts[s] < ends[e]) {
            rooms++;
            s++;
        } else {
            rooms--;
            e++;
        }
        if (rooms > maxRooms) {
            maxRooms = rooms;
        }
    }
    return maxRooms;
}
