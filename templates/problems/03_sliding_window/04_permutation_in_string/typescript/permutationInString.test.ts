import { describe, it, expect } from 'vitest';
import { checkInclusion } from './permutationInString';

describe('PermutationInString', () => {
  it('permutation exists', () => {
    expect(checkInclusion('ab', 'eidbaooo')).toBe(true);
  });

  it('no permutation', () => {
    expect(checkInclusion('ab', 'eidboaoo')).toBe(false);
  });

  it('exact match', () => {
    expect(checkInclusion('abc', 'bca')).toBe(true);
  });

  it('s1 longer than s2', () => {
    expect(checkInclusion('abcdef', 'abc')).toBe(false);
  });

  it('single character match', () => {
    expect(checkInclusion('a', 'a')).toBe(true);
  });

  it('single character no match', () => {
    expect(checkInclusion('a', 'b')).toBe(false);
  });

  it('repeated characters', () => {
    expect(checkInclusion('aab', 'ccccbaa')).toBe(true);
  });
});
