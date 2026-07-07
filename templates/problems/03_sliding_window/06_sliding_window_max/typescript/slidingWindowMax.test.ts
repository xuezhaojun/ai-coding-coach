import { describe, it, expect } from 'vitest';
import { maxSlidingWindow } from './slidingWindowMax';

describe('SlidingWindowMax', () => {
  it('basic case', () => {
    expect(maxSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3)).toEqual([3, 3, 5, 5, 6, 7]);
  });

  it('k equals array length', () => {
    expect(maxSlidingWindow([1, 3, 2], 3)).toEqual([3]);
  });

  it('k equals 1', () => {
    expect(maxSlidingWindow([1, -1, 3], 1)).toEqual([1, -1, 3]);
  });

  it('single element', () => {
    expect(maxSlidingWindow([5], 1)).toEqual([5]);
  });

  it('decreasing sequence', () => {
    expect(maxSlidingWindow([7, 6, 5, 4, 3], 2)).toEqual([7, 6, 5, 4]);
  });

  it('increasing sequence', () => {
    expect(maxSlidingWindow([1, 2, 3, 4, 5], 2)).toEqual([2, 3, 4, 5]);
  });

  it('all same values', () => {
    expect(maxSlidingWindow([4, 4, 4, 4], 2)).toEqual([4, 4, 4]);
  });
});
