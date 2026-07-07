import { describe, it, expect } from 'vitest';
import { containsDuplicate } from './containsDuplicate';

describe('ContainsDuplicate', () => {
    it('has duplicates', () => {
        expect(containsDuplicate([1, 2, 3, 1])).toBe(true);
    });

    it('no duplicates', () => {
        expect(containsDuplicate([1, 2, 3, 4])).toBe(false);
    });

    it('empty slice', () => {
        expect(containsDuplicate([])).toBe(false);
    });

    it('single element', () => {
        expect(containsDuplicate([1])).toBe(false);
    });

    it('all same elements', () => {
        expect(containsDuplicate([5, 5, 5, 5])).toBe(true);
    });

    it('duplicates at end', () => {
        expect(containsDuplicate([1, 2, 3, 4, 5, 5])).toBe(true);
    });

    it('negative numbers with duplicates', () => {
        expect(containsDuplicate([-1, -2, -3, -1])).toBe(true);
    });
});
