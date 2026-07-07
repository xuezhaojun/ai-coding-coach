/**
 * Find the maximum in each sliding window of size k using a monotonic deque.
 * Time: O(n), Space: O(k)
 */
export function solveMaxSlidingWindow(nums: number[], k: number): number[] {
  if (nums.length === 0) {
    return [];
  }

  const deque: number[] = [];
  const result: number[] = [];

  for (let i = 0; i < nums.length; i++) {
    while (deque.length > 0 && deque[0] < i - k + 1) {
      deque.shift();
    }
    while (deque.length > 0 && nums[deque[deque.length - 1]] < nums[i]) {
      deque.pop();
    }
    deque.push(i);
    if (i >= k - 1) {
      result.push(nums[deque[0]]);
    }
  }
  return result;
}
