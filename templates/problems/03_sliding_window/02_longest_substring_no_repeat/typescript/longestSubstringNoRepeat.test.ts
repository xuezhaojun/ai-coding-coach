import { describe, it, expect } from 'vitest';
import { lengthOfLongestSubstring } from './longestSubstringNoRepeat';

describe('LongestSubstringNoRepeat', () => {
  it('basic case', () => {
    expect(lengthOfLongestSubstring('abcabcbb')).toBe(3);
  });

  it('all same characters', () => {
    expect(lengthOfLongestSubstring('bbbbb')).toBe(1);
  });

  it('mixed repeats', () => {
    expect(lengthOfLongestSubstring('pwwkew')).toBe(3);
  });

  it('empty string', () => {
    expect(lengthOfLongestSubstring('')).toBe(0);
  });

  it('single character', () => {
    expect(lengthOfLongestSubstring('a')).toBe(1);
  });

  it('all unique', () => {
    expect(lengthOfLongestSubstring('abcdef')).toBe(6);
  });

  it('spaces and special chars', () => {
    expect(lengthOfLongestSubstring('a b c')).toBe(3);
  });
});
