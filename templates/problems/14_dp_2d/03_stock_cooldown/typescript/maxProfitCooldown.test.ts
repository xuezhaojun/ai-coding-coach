import { describe, it, expect } from 'vitest';
import { maxProfitCooldown } from './maxProfitCooldown';

describe('MaxProfitCooldown', () => {
    it('example 1', () => {
        expect(maxProfitCooldown([1, 2, 3, 0, 2])).toBe(3);
    });

    it('single day', () => {
        expect(maxProfitCooldown([1])).toBe(0);
    });

    it('decreasing prices', () => {
        expect(maxProfitCooldown([5, 4, 3, 2, 1])).toBe(0);
    });

    it('two days profit', () => {
        expect(maxProfitCooldown([1, 2])).toBe(1);
    });

    it('two days no profit', () => {
        expect(maxProfitCooldown([2, 1])).toBe(0);
    });

    it('alternating', () => {
        expect(maxProfitCooldown([1, 4, 2, 7])).toBe(6);
    });
});
