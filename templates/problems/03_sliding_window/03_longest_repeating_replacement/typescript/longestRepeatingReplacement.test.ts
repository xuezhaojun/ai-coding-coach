import { describe, it, expect } from 'vitest';
import { characterReplacement } from './longestRepeatingReplacement';

describe('LongestRepeatingReplacement', () => {
  it('basic case', () => {
    expect(characterReplacement('ABAB', 2)).toBe(4);
  });

  it('replace one', () => {
    expect(characterReplacement('AABABBA', 1)).toBe(4);
  });

  it('no replacement needed', () => {
    expect(characterReplacement('AAAA', 0)).toBe(4);
  });

  it('single character', () => {
    expect(characterReplacement('A', 0)).toBe(1);
  });

  it('k equals string length', () => {
    expect(characterReplacement('ABCD', 4)).toBe(4);
  });

  it('alternating with k=0', () => {
    expect(characterReplacement('ABABAB', 0)).toBe(1);
  });
});
