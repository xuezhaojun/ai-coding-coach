import { describe, it, expect } from 'vitest';
import { change } from './change';

describe('Change', () => {
    it('example 1', () => {
        expect(change(5, [1, 2, 5])).toBe(4);
    });

    it('example 2', () => {
        expect(change(3, [2])).toBe(0);
    });

    it('zero amount', () => {
        expect(change(0, [1, 2])).toBe(1);
    });

    it('single coin', () => {
        expect(change(10, [10])).toBe(1);
    });

    it('single penny', () => {
        expect(change(5, [1])).toBe(1);
    });

    it('two coins', () => {
        expect(change(4, [1, 2])).toBe(3);
    });
});
