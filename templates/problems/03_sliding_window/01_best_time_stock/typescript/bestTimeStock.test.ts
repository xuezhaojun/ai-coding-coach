import { describe, it, expect } from 'vitest';
import { maxProfit } from './bestTimeStock';

describe('BestTimeStock', () => {
  it('basic profit', () => {
    expect(maxProfit([7, 1, 5, 3, 6, 4])).toBe(5);
  });

  it('no profit possible', () => {
    expect(maxProfit([7, 6, 4, 3, 1])).toBe(0);
  });

  it('single day', () => {
    expect(maxProfit([5])).toBe(0);
  });

  it('two days profit', () => {
    expect(maxProfit([1, 2])).toBe(1);
  });

  it('two days no profit', () => {
    expect(maxProfit([2, 1])).toBe(0);
  });

  it('buy at start sell at end', () => {
    expect(maxProfit([1, 2, 3, 4, 5])).toBe(4);
  });

  it('all same price', () => {
    expect(maxProfit([3, 3, 3, 3])).toBe(0);
  });
});
