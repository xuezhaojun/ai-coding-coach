import { describe, it, expect } from 'vitest';
import { plusOne } from './plusOne';

describe('PlusOne', () => {
    it('no carry', () => {
        expect(plusOne([1, 2, 3])).toEqual([1, 2, 4]);
    });

    it('single carry', () => {
        expect(plusOne([4, 3, 2, 9])).toEqual([4, 3, 3, 0]);
    });

    it('all nines', () => {
        expect(plusOne([9, 9, 9])).toEqual([1, 0, 0, 0]);
    });

    it('single digit', () => {
        expect(plusOne([0])).toEqual([1]);
    });

    it('single nine', () => {
        expect(plusOne([9])).toEqual([1, 0]);
    });

    it('large number', () => {
        expect(plusOne([8, 9, 9, 9])).toEqual([9, 0, 0, 0]);
    });
});
