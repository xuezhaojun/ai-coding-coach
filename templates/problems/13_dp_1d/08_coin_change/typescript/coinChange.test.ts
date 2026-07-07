import { describe, it, expect } from 'vitest';
import { coinChange } from './coinChange';

describe('CoinChange', () => {
    it('example 1', () => {
        expect(coinChange([1, 2, 5], 11)).toBe(3);
    });

    it('impossible', () => {
        expect(coinChange([2], 3)).toBe(-1);
    });

    it('zero amount', () => {
        expect(coinChange([1], 0)).toBe(0);
    });

    it('single coin exact', () => {
        expect(coinChange([1], 1)).toBe(1);
    });

    it('single coin multiple', () => {
        expect(coinChange([1], 5)).toBe(5);
    });

    it('large coins small amount', () => {
        expect(coinChange([5, 10], 3)).toBe(-1);
    });

    it('multiple denominations', () => {
        expect(coinChange([1, 5, 10, 25], 30)).toBe(2);
    });
});
