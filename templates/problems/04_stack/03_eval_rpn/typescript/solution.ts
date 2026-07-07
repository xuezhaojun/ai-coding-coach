/**
 * Evaluate a reverse polish notation expression using a stack.
 * Time: O(n), Space: O(n)
 */
export function solveEvalRPN(tokens: string[]): number {
  const stack: number[] = [];
  for (const token of tokens) {
    if (token === '+' || token === '-' || token === '*' || token === '/') {
      const b = stack.pop()!;
      const a = stack.pop()!;
      let result: number;
      switch (token) {
        case '+':
          result = a + b;
          break;
        case '-':
          result = a - b;
          break;
        case '*':
          result = a * b;
          break;
        default:
          // Integer division truncating toward zero (matches Go behavior)
          result = Math.trunc(a / b);
          break;
      }
      stack.push(result);
    } else {
      stack.push(parseInt(token, 10));
    }
  }
  return stack[0];
}
