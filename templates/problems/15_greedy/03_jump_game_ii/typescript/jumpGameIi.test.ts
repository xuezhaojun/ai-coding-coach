import { describe, it, expect } from 'vitest';
import { jump } from './jumpGameIi';

describe('JumpGameII', () => {
    it('example 1', () => {
        expect(jump([2, 3, 1, 1, 4])).toBe(2);
    });

    it('example 2', () => {
        expect(jump([2, 3, 0, 1, 4])).toBe(2);
    });

    it('single element', () => {
        expect(jump([0])).toBe(0);
    });

    it('two elements', () => {
        expect(jump([1, 2])).toBe(1);
    });

    it('already at end', () => {
        expect(jump([1])).toBe(0);
    });

    it('linear jumps', () => {
        expect(jump([1, 1, 1, 1])).toBe(3);
    });

    it('large first jump', () => {
        expect(jump([5, 1, 1, 1, 1, 1])).toBe(1);
    });

    it('greedy per level', () => {
        expect(jump([1, 2, 1, 1, 1])).toBe(3);
    });

    it('level boundary', () => {
        expect(jump([7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3])).toBe(2);
    });
});
