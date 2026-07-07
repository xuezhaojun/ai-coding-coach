import { describe, it, expect } from 'vitest';
import { missingNumber } from './missingNumber';

describe('MissingNumber', () => {
    it('missing 2', () => {
        expect(missingNumber([3, 0, 1])).toBe(2);
    });

    it('missing 2 from 3', () => {
        expect(missingNumber([0, 1])).toBe(2);
    });

    it('missing 8', () => {
        expect(missingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1])).toBe(8);
    });

    it('missing 0', () => {
        expect(missingNumber([1])).toBe(0);
    });

    it('missing 1', () => {
        expect(missingNumber([0])).toBe(1);
    });

    it('single zero', () => {
        expect(missingNumber([0, 2, 3])).toBe(1);
    });
});
