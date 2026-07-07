/**
 * Group strings that are anagrams using sorted-key hashing.
 * Time: O(n * k log k) where k is max string length, Space: O(n * k)
 */
export function solveGroupAnagrams(strs: string[]): string[][] {
    const groups = new Map<string, string[]>();
    for (const s of strs) {
        const key = new Array(26).fill(0);
        for (let i = 0; i < s.length; i++) {
            key[s.charCodeAt(i) - 'a'.charCodeAt(0)]++;
        }
        const keyStr = key.join(',');
        if (!groups.has(keyStr)) {
            groups.set(keyStr, []);
        }
        groups.get(keyStr)!.push(s);
    }
    return Array.from(groups.values());
}
