import { describe, it, expect } from 'vitest';
import { productExceptSelf } from './productExceptSelf';

describe('ProductExceptSelf', () => {
    it('basic case', () => {
        expect(productExceptSelf([1, 2, 3, 4])).toEqual([24, 12, 8, 6]);
    });

    it('contains zero', () => {
        expect(productExceptSelf([-1, 1, 0, -3, 3])).toEqual([0, 0, 9, 0, 0]);
    });

    it('two elements', () => {
        expect(productExceptSelf([2, 3])).toEqual([3, 2]);
    });

    it('all ones', () => {
        expect(productExceptSelf([1, 1, 1, 1])).toEqual([1, 1, 1, 1]);
    });

    it('contains negative numbers', () => {
        expect(productExceptSelf([-1, 2, -3, 4])).toEqual([-24, 12, -8, 6]);
    });

    it('two zeros', () => {
        expect(productExceptSelf([0, 0, 1, 2])).toEqual([0, 0, 0, 0]);
    });

    it('large values', () => {
        expect(productExceptSelf([10, 20, 30])).toEqual([600, 300, 200]);
    });
});
