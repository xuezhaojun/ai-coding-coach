import { describe, it, expect } from 'vitest';
import { isNStraightHand } from './handOfStraights';

describe('HandOfStraights', () => {
    it('example true', () => {
        expect(isNStraightHand([1, 2, 3, 6, 2, 3, 4, 7, 8], 3)).toBe(true);
    });

    it('example false', () => {
        expect(isNStraightHand([1, 2, 3, 4, 5], 4)).toBe(false);
    });

    it('single group', () => {
        expect(isNStraightHand([1, 2, 3], 3)).toBe(true);
    });

    it('group size 1', () => {
        expect(isNStraightHand([5, 3, 1], 1)).toBe(true);
    });

    it('not divisible', () => {
        expect(isNStraightHand([1, 2, 3, 4], 3)).toBe(false);
    });

    it('duplicates needed', () => {
        expect(isNStraightHand([1, 1, 2, 2, 3, 3], 3)).toBe(true);
    });

    it('gap in sequence', () => {
        expect(isNStraightHand([1, 2, 4, 5, 6, 7], 3)).toBe(false);
    });
});
