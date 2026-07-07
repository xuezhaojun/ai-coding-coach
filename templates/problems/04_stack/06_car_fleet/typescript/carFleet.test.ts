import { describe, it, expect } from 'vitest';
import { carFleet } from './carFleet';

describe('CarFleet', () => {
  it('example 1', () => {
    expect(carFleet(12, [10, 8, 0, 5, 3], [2, 4, 1, 1, 3])).toBe(3);
  });

  it('single car', () => {
    expect(carFleet(10, [3], [3])).toBe(1);
  });

  it('all same speed', () => {
    expect(carFleet(10, [0, 2, 4], [2, 2, 2])).toBe(3);
  });

  it('all merge', () => {
    expect(carFleet(10, [0, 2], [2, 1])).toBe(1);
  });

  it('no cars', () => {
    expect(carFleet(10, [], [])).toBe(0);
  });

  it('two separate fleets', () => {
    expect(carFleet(100, [0, 50], [1, 1])).toBe(2);
  });
});
