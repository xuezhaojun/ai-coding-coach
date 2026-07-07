import { describe, it, expect } from 'vitest';
import { evalRPN } from './evalRpn';

describe('EvalRpn', () => {
  it('simple addition', () => {
    expect(evalRPN(['2', '1', '+'])).toBe(3);
  });

  it('simple subtraction', () => {
    expect(evalRPN(['4', '2', '-'])).toBe(2);
  });

  it('simple multiplication', () => {
    expect(evalRPN(['3', '4', '*'])).toBe(12);
  });

  it('simple division', () => {
    expect(evalRPN(['6', '3', '/'])).toBe(2);
  });

  it('complex expression', () => {
    expect(evalRPN(['2', '1', '+', '3', '*'])).toBe(9);
  });

  it('leetcode example', () => {
    expect(evalRPN(['4', '13', '5', '/', '+'])).toBe(6);
  });

  it('negative result', () => {
    expect(evalRPN(['1', '2', '-'])).toBe(-1);
  });

  it('single number', () => {
    expect(evalRPN(['42'])).toBe(42);
  });
});
