import { describe, it, expect } from 'vitest';
import { generateParenthesis } from './generateParentheses';

describe('GenerateParentheses', () => {
  it('n=1', () => {
    expect(generateParenthesis(1).sort()).toEqual(['()']);
  });

  it('n=2', () => {
    expect(generateParenthesis(2).sort()).toEqual(['(())', '()()']);
  });

  it('n=3', () => {
    expect(generateParenthesis(3).sort()).toEqual([
      '((()))',
      '(()())',
      '(())()',
      '()(())',
      '()()()',
    ]);
  });

  it('n=4 count', () => {
    expect(generateParenthesis(4).length).toBe(14);
  });

  it('n=0', () => {
    expect(generateParenthesis(0)).toEqual(['']);
  });
});
