import { describe, it, expect } from 'vitest';
import { getSum } from './sumTwoIntegers';

describe('SumTwoIntegers', () => {
    it('both positive', () => {
        expect(getSum(1, 2)).toBe(3);
    });

    it('positive and negative', () => {
        expect(getSum(2, -1)).toBe(1);
    });

    it('both negative', () => {
        expect(getSum(-1, -1)).toBe(-2);
    });

    it('zero and number', () => {
        expect(getSum(0, 5)).toBe(5);
    });

    it('both zero', () => {
        expect(getSum(0, 0)).toBe(0);
    });

    it('larger numbers', () => {
        expect(getSum(100, 200)).toBe(300);
    });

    it('negative result', () => {
        expect(getSum(-5, 3)).toBe(-2);
    });
});
