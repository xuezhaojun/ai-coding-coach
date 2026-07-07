import { describe, it, expect } from 'vitest';
import { climbStairs } from './climbStairs';

describe('ClimbStairs', () => {
    it('one step', () => {
        expect(climbStairs(1)).toBe(1);
    });

    it('two steps', () => {
        expect(climbStairs(2)).toBe(2);
    });

    it('three steps', () => {
        expect(climbStairs(3)).toBe(3);
    });

    it('four steps', () => {
        expect(climbStairs(4)).toBe(5);
    });

    it('five steps', () => {
        expect(climbStairs(5)).toBe(8);
    });

    it('ten steps', () => {
        expect(climbStairs(10)).toBe(89);
    });
});
