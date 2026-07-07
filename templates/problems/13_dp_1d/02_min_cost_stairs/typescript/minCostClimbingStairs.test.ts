import { describe, it, expect } from 'vitest';
import { minCostClimbingStairs } from './minCostClimbingStairs';

describe('MinCostClimbingStairs', () => {
    it('example 1', () => {
        expect(minCostClimbingStairs([10, 15, 20])).toBe(15);
    });

    it('example 2', () => {
        expect(minCostClimbingStairs([1, 100, 1, 1, 1, 100, 1, 1, 100, 1])).toBe(6);
    });

    it('two steps equal cost', () => {
        expect(minCostClimbingStairs([5, 5])).toBe(5);
    });

    it('two steps different cost', () => {
        expect(minCostClimbingStairs([1, 100])).toBe(1);
    });

    it('increasing cost', () => {
        expect(minCostClimbingStairs([1, 2, 3, 4, 5])).toBe(6);
    });

    it('all zeros', () => {
        expect(minCostClimbingStairs([0, 0, 0, 0])).toBe(0);
    });
});
