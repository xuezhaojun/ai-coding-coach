import { describe, it, expect } from 'vitest';
import { minWindow } from './minWindowSubstring';

describe('MinWindowSubstring', () => {
  it('basic case', () => {
    expect(minWindow('ADOBECODEBANC', 'ABC')).toBe('BANC');
  });

  it('exact match', () => {
    expect(minWindow('a', 'a')).toBe('a');
  });

  it('no valid window', () => {
    expect(minWindow('a', 'aa')).toBe('');
  });

  it('t not in s', () => {
    expect(minWindow('abc', 'z')).toBe('');
  });

  it('entire string is window', () => {
    expect(minWindow('abc', 'abc')).toBe('abc');
  });

  it('duplicate chars in t', () => {
    expect(minWindow('aaabbc', 'aab')).toBe('aab');
  });

  it('empty s', () => {
    expect(minWindow('', 'a')).toBe('');
  });

  it('many duplicates before target chars', () => {
    expect(minWindow('aaaaaaaaaaaabbbbbcdd', 'abcdd')).toBe('abbbbbcdd');
  });
});
