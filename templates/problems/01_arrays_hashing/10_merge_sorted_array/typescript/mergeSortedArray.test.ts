import { describe, it, expect } from 'vitest';
import { merge } from './mergeSortedArray';

describe('MergeSortedArray', () => {
    it('basic merge', () => {
        const nums1 = [1, 2, 3, 0, 0, 0];
        merge(nums1, 3, [2, 5, 6], 3);
        expect(nums1).toEqual([1, 2, 2, 3, 5, 6]);
    });

    it('nums2 empty', () => {
        const nums1 = [1];
        merge(nums1, 1, [], 0);
        expect(nums1).toEqual([1]);
    });

    it('nums1 empty', () => {
        const nums1 = [0];
        merge(nums1, 0, [1], 1);
        expect(nums1).toEqual([1]);
    });

    it('all nums1 smaller', () => {
        const nums1 = [1, 2, 3, 0, 0, 0];
        merge(nums1, 3, [4, 5, 6], 3);
        expect(nums1).toEqual([1, 2, 3, 4, 5, 6]);
    });

    it('all nums2 smaller', () => {
        const nums1 = [4, 5, 6, 0, 0, 0];
        merge(nums1, 3, [1, 2, 3], 3);
        expect(nums1).toEqual([1, 2, 3, 4, 5, 6]);
    });

    it('duplicates and interleaving', () => {
        const nums1 = [1, 2, 4, 5, 6, 0, 0, 0];
        merge(nums1, 5, [2, 3, 7], 3);
        expect(nums1).toEqual([1, 2, 2, 3, 4, 5, 6, 7]);
    });
});
