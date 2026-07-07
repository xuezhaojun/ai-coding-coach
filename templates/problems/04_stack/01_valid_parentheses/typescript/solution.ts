/**
 * Check if brackets are valid using a stack.
 * Time: O(n), Space: O(n)
 */
export function solveIsValid(s: string): boolean {
  const stack: string[] = [];
  const pairs: Record<string, string> = { ')': '(', ']': '[', '}': '{' };
  for (const c of s) {
    if (c === '(' || c === '[' || c === '{') {
      stack.push(c);
    } else {
      if (stack.length === 0 || stack[stack.length - 1] !== pairs[c]) {
        return false;
      }
      stack.pop();
    }
  }
  return stack.length === 0;
}
