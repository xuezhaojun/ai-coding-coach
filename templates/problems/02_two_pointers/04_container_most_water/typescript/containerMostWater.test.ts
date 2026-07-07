import { describe, it, expect } from 'vitest';
import { maxArea } from './containerMostWater';

describe('ContainerMostWater', () => {
    it('basic case', () => {
        expect(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7])).toBe(49);
    });

    it('two elements', () => {
        expect(maxArea([1, 1])).toBe(1);
    });

    it('decreasing heights', () => {
        expect(maxArea([4, 3, 2, 1, 4])).toBe(16);
    });

    it('equal heights', () => {
        expect(maxArea([5, 5, 5, 5])).toBe(15);
    });

    it('one tall wall', () => {
        expect(maxArea([1, 2, 1])).toBe(2);
    });

    it('large difference', () => {
        expect(maxArea([1, 1000, 1000, 1])).toBe(1000);
    });
});
