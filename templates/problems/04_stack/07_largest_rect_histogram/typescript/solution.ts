/**
 * Use a monotonic increasing stack.
 * Time: O(n), Space: O(n)
 */
export function solveLargestRectangleArea(heights: number[]): number {
  const stack: number[] = []; // stack of indices
  let maxArea = 0;
  for (let i = 0; i <= heights.length; i++) {
    const h = i < heights.length ? heights[i] : 0;
    while (stack.length > 0 && h < heights[stack[stack.length - 1]]) {
      const height = heights[stack.pop()!];
      const width = stack.length === 0 ? i : i - stack[stack.length - 1] - 1;
      const area = height * width;
      if (area > maxArea) {
        maxArea = area;
      }
    }
    stack.push(i);
  }
  return maxArea;
}
