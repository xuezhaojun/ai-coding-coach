import { describe, it, expect } from 'vitest';
import { isAnagram } from './validAnagram';

describe('ValidAnagram', () => {
    it('valid anagram', () => {
        expect(isAnagram('anagram', 'nagaram')).toBe(true);
    });

    it('not anagram', () => {
        expect(isAnagram('rat', 'car')).toBe(false);
    });

    it('empty strings', () => {
        expect(isAnagram('', '')).toBe(true);
    });

    it('different lengths', () => {
        expect(isAnagram('ab', 'abc')).toBe(false);
    });

    it('single characters same', () => {
        expect(isAnagram('a', 'a')).toBe(true);
    });

    it('single characters different', () => {
        expect(isAnagram('a', 'b')).toBe(false);
    });

    it('repeated characters', () => {
        expect(isAnagram('aabb', 'bbaa')).toBe(true);
    });

    it('same chars different frequency', () => {
        expect(isAnagram('aaab', 'aabb')).toBe(false);
    });
});
