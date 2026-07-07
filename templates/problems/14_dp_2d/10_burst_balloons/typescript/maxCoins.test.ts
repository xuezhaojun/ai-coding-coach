import { describe, it, expect } from 'vitest';
import { maxCoins } from './maxCoins';

describe('MaxCoins', () => {
    it('example 1', () => {
        expect(maxCoins([3, 1, 5, 8])).toBe(167);
    });

    it('single balloon', () => {
        expect(maxCoins([5])).toBe(5);
    });

    it('two balloons', () => {
        expect(maxCoins([1, 5])).toBe(10);
    });

    it('example 2', () => {
        expect(maxCoins([1, 5])).toBe(10);
    });

    it('all ones', () => {
        expect(maxCoins([1, 1, 1])).toBe(3);
    });

    it('descending', () => {
        expect(maxCoins([5, 4, 3, 2, 1])).toBe(110);
    });
});
