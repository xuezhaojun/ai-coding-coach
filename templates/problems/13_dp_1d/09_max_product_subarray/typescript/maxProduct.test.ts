import { describe, it, expect } from 'vitest';
import { maxProduct } from './maxProduct';

describe('MaxProduct', () => {
    it('example 1', () => {
        expect(maxProduct([2, 3, -2, 4])).toBe(6);
    });

    it('example 2', () => {
        expect(maxProduct([-2, 0, -1])).toBe(0);
    });

    it('single negative', () => {
        expect(maxProduct([-2])).toBe(-2);
    });

    it('two negatives', () => {
        expect(maxProduct([-2, -3])).toBe(6);
    });

    it('contains zero', () => {
        expect(maxProduct([-2, 3, -4])).toBe(24);
    });

    it('all positive', () => {
        expect(maxProduct([1, 2, 3, 4])).toBe(24);
    });

    it('single element', () => {
        expect(maxProduct([5])).toBe(5);
    });
});
