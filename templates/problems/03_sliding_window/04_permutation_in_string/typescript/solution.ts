/**
 * Check if any permutation of s1 is a substring of s2.
 * Time: O(n), Space: O(1) - fixed 26-letter alphabet
 */
export function solveCheckInclusion(s1: string, s2: string): boolean {
  if (s1.length > s2.length) {
    return false;
  }

  const s1Count: number[] = new Array(26).fill(0);
  const windowCount: number[] = new Array(26).fill(0);
  for (let i = 0; i < s1.length; i++) {
    s1Count[s1.charCodeAt(i) - 97]++; // 'a' = 97
    windowCount[s2.charCodeAt(i) - 97]++;
  }

  if (arraysEqual(s1Count, windowCount)) {
    return true;
  }

  for (let i = s1.length; i < s2.length; i++) {
    windowCount[s2.charCodeAt(i) - 97]++;
    windowCount[s2.charCodeAt(i - s1.length) - 97]--;
    if (arraysEqual(s1Count, windowCount)) {
      return true;
    }
  }
  return false;
}

function arraysEqual(a: number[], b: number[]): boolean {
  for (let i = 0; i < a.length; i++) {
    if (a[i] !== b[i]) return false;
  }
  return true;
}
