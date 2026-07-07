import { describe, it, expect } from 'vitest';
import { isPalindrome } from './validPalindrome';

describe('ValidPalindrome', () => {
    it('alphanumeric palindrome with spaces and punctuation', () => {
        expect(isPalindrome('A man, a plan, a canal: Panama')).toBe(true);
    });

    it('not a palindrome', () => {
        expect(isPalindrome('race a car')).toBe(false);
    });

    it('empty string is palindrome', () => {
        expect(isPalindrome(' ')).toBe(true);
    });

    it('single character', () => {
        expect(isPalindrome('a')).toBe(true);
    });

    it('mixed case palindrome', () => {
        expect(isPalindrome('Aa')).toBe(true);
    });

    it('digits in palindrome', () => {
        expect(isPalindrome('0P')).toBe(false);
    });
});
