import { describe, it, expect } from 'vitest';
import { isValid } from './validParentheses';

describe('ValidParentheses', () => {
  it('empty string', () => {
    expect(isValid('')).toBe(true);
  });

  it('single pair parens', () => {
    expect(isValid('()')).toBe(true);
  });

  it('multiple types', () => {
    expect(isValid('()[]{}')).toBe(true);
  });

  it('nested', () => {
    expect(isValid('{[]}')).toBe(true);
  });

  it('mismatched', () => {
    expect(isValid('(]')).toBe(false);
  });

  it('unmatched open', () => {
    expect(isValid('([')).toBe(false);
  });

  it('complex valid', () => {
    expect(isValid('({[()]})')).toBe(true);
  });

  it('single char', () => {
    expect(isValid('(')).toBe(false);
  });

  it('unmatched close', () => {
    expect(isValid(']')).toBe(false);
  });

  it('leading close', () => {
    expect(isValid(']()')).toBe(false);
  });
});
