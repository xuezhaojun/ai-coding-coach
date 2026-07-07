import { describe, it, expect } from 'vitest';
import { wordBreak } from './wordBreak';

describe('WordBreak', () => {
    it('example leetcode', () => {
        expect(wordBreak('leetcode', ['leet', 'code'])).toBe(true);
    });

    it('example applepenapple', () => {
        expect(wordBreak('applepenapple', ['apple', 'pen'])).toBe(true);
    });

    it('cannot break', () => {
        expect(wordBreak('catsandog', ['cats', 'dog', 'sand', 'and', 'cat'])).toBe(false);
    });

    it('empty string', () => {
        expect(wordBreak('', ['a'])).toBe(true);
    });

    it('single char match', () => {
        expect(wordBreak('a', ['a'])).toBe(true);
    });

    it('single char no match', () => {
        expect(wordBreak('b', ['a'])).toBe(false);
    });

    it('overlapping words', () => {
        expect(wordBreak('cars', ['car', 'ca', 'rs'])).toBe(true);
    });
});
