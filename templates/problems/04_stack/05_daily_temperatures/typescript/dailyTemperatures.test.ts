import { describe, it, expect } from 'vitest';
import { dailyTemperatures } from './dailyTemperatures';

describe('DailyTemperatures', () => {
  it('typical case', () => {
    expect(dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73])).toEqual([1, 1, 4, 2, 1, 1, 0, 0]);
  });

  it('decreasing', () => {
    expect(dailyTemperatures([30, 20, 10])).toEqual([0, 0, 0]);
  });

  it('increasing', () => {
    expect(dailyTemperatures([10, 20, 30])).toEqual([1, 1, 0]);
  });

  it('single element', () => {
    expect(dailyTemperatures([50])).toEqual([0]);
  });

  it('all same', () => {
    expect(dailyTemperatures([70, 70, 70, 70])).toEqual([0, 0, 0, 0]);
  });

  it('two elements warmer', () => {
    expect(dailyTemperatures([30, 60])).toEqual([1, 0]);
  });

  it('two elements cooler', () => {
    expect(dailyTemperatures([60, 30])).toEqual([0, 0]);
  });
});
