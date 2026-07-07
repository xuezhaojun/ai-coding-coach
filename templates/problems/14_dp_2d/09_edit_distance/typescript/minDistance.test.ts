import { describe, it, expect } from 'vitest';
import { minDistance } from './minDistance';

describe('MinDistance', () => {
    it('horse to ros', () => {
        expect(minDistance('horse', 'ros')).toBe(3);
    });

    it('intention to execution', () => {
        expect(minDistance('intention', 'execution')).toBe(5);
    });

    it('empty to abc', () => {
        expect(minDistance('', 'abc')).toBe(3);
    });

    it('abc to empty', () => {
        expect(minDistance('abc', '')).toBe(3);
    });

    it('both empty', () => {
        expect(minDistance('', '')).toBe(0);
    });

    it('same strings', () => {
        expect(minDistance('abc', 'abc')).toBe(0);
    });

    it('single char different', () => {
        expect(minDistance('a', 'b')).toBe(1);
    });
});
