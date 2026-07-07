/**
 * Find the longest substring without repeating characters.
 * Time: O(n), Space: O(min(n, 128)) - ASCII character set
 */
export function solveLengthOfLongestSubstring(s: string): number {
  const lastSeen = new Map<string, number>();
  let best = 0;
  let left = 0;

  for (let right = 0; right < s.length; right++) {
    const ch = s[right];
    if (lastSeen.has(ch) && (lastSeen.get(ch) as number) >= left) {
      left = (lastSeen.get(ch) as number) + 1;
    }
    if (right - left + 1 > best) {
      best = right - left + 1;
    }
    lastSeen.set(ch, right);
  }
  return best;
}
