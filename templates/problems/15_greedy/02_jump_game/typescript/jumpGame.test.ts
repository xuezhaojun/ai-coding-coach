import { describe, it, expect } from 'vitest';
import { canJump } from './jumpGame';

describe('JumpGame', () => {
    it('reachable', () => {
        expect(canJump([2, 3, 1, 1, 4])).toBe(true);
    });

    it('not reachable', () => {
        expect(canJump([3, 2, 1, 0, 4])).toBe(false);
    });

    it('single element', () => {
        expect(canJump([0])).toBe(true);
    });

    it('two elements reachable', () => {
        expect(canJump([1, 0])).toBe(true);
    });

    it('all zeros except first', () => {
        expect(canJump([5, 0, 0, 0, 0])).toBe(true);
    });

    it('large first jump', () => {
        expect(canJump([10, 0, 0, 0, 0, 0, 0, 0, 1])).toBe(true);
    });

    it('stuck at zero', () => {
        expect(canJump([1, 0, 1])).toBe(false);
    });
});
