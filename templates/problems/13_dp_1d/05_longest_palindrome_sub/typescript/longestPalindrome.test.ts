import { describe, it, expect } from 'vitest';
import { longestPalindrome } from './longestPalindrome';

describe('LongestPalindrome', () => {
    it.each([
        { name: 'example babad', s: 'babad', valid: ['bab', 'aba'] },
        { name: 'example cbbd', s: 'cbbd', valid: ['bb'] },
        { name: 'single character', s: 'a', valid: ['a'] },
        { name: 'all same characters', s: 'aaaa', valid: ['aaaa'] },
        { name: 'entire string is palindrome', s: 'racecar', valid: ['racecar'] },
        { name: 'no repeats', s: 'abcde', valid: ['a', 'b', 'c', 'd', 'e'] },
        { name: 'two characters same', s: 'bb', valid: ['bb'] },
    ])('case: $name', ({ s, valid }) => {
        const result = longestPalindrome(s);
        expect(valid).toContain(result);
    });
});
