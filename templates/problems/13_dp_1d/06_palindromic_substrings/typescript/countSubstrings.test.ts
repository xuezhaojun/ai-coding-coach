import { describe, it, expect } from 'vitest';
import { countSubstrings } from './countSubstrings';

describe('CountSubstrings', () => {
    it('example abc', () => {
        expect(countSubstrings('abc')).toBe(3);
    });

    it('example aaa', () => {
        expect(countSubstrings('aaa')).toBe(6);
    });

    it('single char', () => {
        expect(countSubstrings('a')).toBe(1);
    });

    it('two same chars', () => {
        expect(countSubstrings('aa')).toBe(3);
    });

    it('two different chars', () => {
        expect(countSubstrings('ab')).toBe(2);
    });

    it('racecar', () => {
        expect(countSubstrings('racecar')).toBe(10);
    });
});
