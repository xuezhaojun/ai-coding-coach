import { describe, it, expect } from 'vitest';
import { singleNumber } from './singleNumber';

describe('SingleNumber', () => {
    it('example 1', () => {
        expect(singleNumber([2, 2, 1])).toBe(1);
    });

    it('example 2', () => {
        expect(singleNumber([4, 1, 2, 1, 2])).toBe(4);
    });

    it('single element', () => {
        expect(singleNumber([1])).toBe(1);
    });

    it('negative numbers', () => {
        expect(singleNumber([-1, -1, -2])).toBe(-2);
    });

    it('zero is single', () => {
        expect(singleNumber([0, 1, 1])).toBe(0);
    });

    it('larger array', () => {
        expect(singleNumber([1, 3, 1, 2, 3])).toBe(2);
    });
});
