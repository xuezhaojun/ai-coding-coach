import { describe, it, expect } from 'vitest';
import { largestRectangleArea } from './largestRectHistogram';

describe('LargestRectHistogram', () => {
  it('example 1', () => {
    expect(largestRectangleArea([2, 1, 5, 6, 2, 3])).toBe(10);
  });

  it('single bar', () => {
    expect(largestRectangleArea([5])).toBe(5);
  });

  it('increasing', () => {
    expect(largestRectangleArea([1, 2, 3, 4, 5])).toBe(9);
  });

  it('decreasing', () => {
    expect(largestRectangleArea([5, 4, 3, 2, 1])).toBe(9);
  });

  it('all same', () => {
    expect(largestRectangleArea([3, 3, 3, 3])).toBe(12);
  });

  it('two bars', () => {
    expect(largestRectangleArea([2, 4])).toBe(4);
  });

  it('valley', () => {
    expect(largestRectangleArea([6, 2, 5, 4, 5, 1, 6])).toBe(12);
  });

  it('empty', () => {
    expect(largestRectangleArea([])).toBe(0);
  });
});
