import { describe, it, expect } from 'vitest';
import { MinStack } from './minStack';

function runOperations(
  operations: string[],
  values: number[],
  expected: number[]
): void {
  const s = new MinStack();
  for (let i = 0; i < operations.length; i++) {
    const op = operations[i];
    if (op === 'Push') {
      s.push(values[i]);
    } else if (op === 'Pop') {
      s.pop();
    } else if (op === 'Top') {
      expect(s.top()).toBe(expected[i]);
    } else if (op === 'GetMin') {
      expect(s.getMin()).toBe(expected[i]);
    }
  }
}

describe('MinStack', () => {
  it('basic push and get min', () => {
    runOperations(
      ['Push', 'Push', 'Push', 'GetMin', 'Pop', 'Top', 'GetMin'],
      [-2, 0, -3, 0, 0, 0, 0],
      [0, 0, 0, -3, 0, 0, -2]
    );
  });

  it('single element', () => {
    runOperations(['Push', 'Top', 'GetMin'], [5, 0, 0], [0, 5, 5]);
  });

  it('decreasing order', () => {
    runOperations(
      ['Push', 'Push', 'Push', 'GetMin', 'Pop', 'GetMin'],
      [3, 2, 1, 0, 0, 0],
      [0, 0, 0, 1, 0, 2]
    );
  });

  it('increasing order', () => {
    runOperations(
      ['Push', 'Push', 'Push', 'GetMin', 'Pop', 'GetMin'],
      [1, 2, 3, 0, 0, 0],
      [0, 0, 0, 1, 0, 1]
    );
  });

  it('duplicate minimums', () => {
    runOperations(
      ['Push', 'Push', 'Push', 'GetMin', 'Pop', 'GetMin'],
      [1, 1, 1, 0, 0, 0],
      [0, 0, 0, 1, 0, 1]
    );
  });
});
