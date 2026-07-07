/**
 * Find the minimum window in s that contains all characters of t.
 * Time: O(n), Space: O(1) - fixed character set
 */
export function solveMinWindow(s: string, t: string): string {
  if (t.length === 0) {
    return '';
  }

  const tCount: number[] = new Array(128).fill(0);
  const windowCount: number[] = new Array(128).fill(0);
  for (let i = 0; i < t.length; i++) {
    tCount[t.charCodeAt(i)]++;
  }

  let need = 0;
  for (let i = 0; i < 128; i++) {
    if (tCount[i] > 0) need++;
  }

  let have = 0;
  let bestLen = s.length + 1;
  let bestStart = 0;
  let left = 0;

  for (let right = 0; right < s.length; right++) {
    const c = s.charCodeAt(right);
    windowCount[c]++;
    if (tCount[c] > 0 && windowCount[c] === tCount[c]) {
      have++;
    }

    while (have === need) {
      if (right - left + 1 < bestLen) {
        bestLen = right - left + 1;
        bestStart = left;
      }
      const lc = s.charCodeAt(left);
      windowCount[lc]--;
      if (tCount[lc] > 0 && windowCount[lc] < tCount[lc]) {
        have--;
      }
      left++;
    }
  }

  if (bestLen > s.length) {
    return '';
  }
  return s.substring(bestStart, bestStart + bestLen);
}
