/**
 * Find all palindrome partitions of the string.
 * Time: O(n * 2^n), Space: O(n) recursion depth
 */
export function solvePartition(s: string): string[][] {
    const result: string[][] = [];

    const isPalin = (sub: string): boolean => {
        let l = 0, r = sub.length - 1;
        while (l < r) {
            if (sub[l] !== sub[r]) return false;
            l++;
            r--;
        }
        return true;
    };

    const backtrack = (start: number, current: string[]): void => {
        if (start === s.length) {
            result.push([...current]);
            return;
        }
        for (let end = start + 1; end <= s.length; end++) {
            const sub = s.slice(start, end);
            if (isPalin(sub)) {
                current.push(sub);
                backtrack(end, current);
                current.pop();
            }
        }
    };

    backtrack(0, []);
    return result;
}
