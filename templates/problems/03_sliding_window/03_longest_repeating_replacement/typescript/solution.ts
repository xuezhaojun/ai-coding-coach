/**
 * Find the longest substring with at most k replacements.
 * Time: O(n), Space: O(1) - fixed 26-letter alphabet
 */
export function solveCharacterReplacement(s: string, k: number): number {
  const count: number[] = new Array(26).fill(0);
  let maxFreq = 0;
  let left = 0;
  let best = 0;

  for (let right = 0; right < s.length; right++) {
    const idx = s.charCodeAt(right) - 65; // 'A' = 65
    count[idx]++;
    if (count[idx] > maxFreq) {
      maxFreq = count[idx];
    }
    while (right - left + 1 - maxFreq > k) {
      count[s.charCodeAt(left) - 65]--;
      left++;
    }
    if (right - left + 1 > best) {
      best = right - left + 1;
    }
  }
  return best;
}
