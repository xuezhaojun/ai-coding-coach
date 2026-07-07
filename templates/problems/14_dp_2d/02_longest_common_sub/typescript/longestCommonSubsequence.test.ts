import { describe, it, expect } from 'vitest';
import { longestCommonSubsequence } from './longestCommonSubsequence';

describe('LongestCommonSubsequence', () => {
    it('example 1', () => {
        expect(longestCommonSubsequence('abcde', 'ace')).toBe(3);
    });

    it('example 2', () => {
        expect(longestCommonSubsequence('abc', 'abc')).toBe(3);
    });

    it('no common', () => {
        expect(longestCommonSubsequence('abc', 'def')).toBe(0);
    });

    it('one char vs longer', () => {
        expect(longestCommonSubsequence('a', 'abc')).toBe(1);
    });

    it('single char match', () => {
        expect(longestCommonSubsequence('a', 'a')).toBe(1);
    });

    it('single char no match', () => {
        expect(longestCommonSubsequence('a', 'b')).toBe(0);
    });

    it('longer example', () => {
        expect(longestCommonSubsequence('oxcpqrsvwf', 'shmtulqrypy')).toBe(2);
    });

    it('duplicate chars must not double count', () => {
        expect(longestCommonSubsequence('bsbininm', 'jmjkbkjkv')).toBe(1);
    });
});
