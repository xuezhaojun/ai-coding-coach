import { describe, it, expect } from 'vitest';
import { partitionLabels } from './partitionLabels';

describe('PartitionLabels', () => {
    it('example 1', () => {
        expect(partitionLabels('ababcbacadefegdehijhklij')).toEqual([9, 7, 8]);
    });

    it('single char', () => {
        expect(partitionLabels('a')).toEqual([1]);
    });

    it('all same', () => {
        expect(partitionLabels('aaaa')).toEqual([4]);
    });

    it('all unique', () => {
        expect(partitionLabels('abcdef')).toEqual([1, 1, 1, 1, 1, 1]);
    });

    it('two partitions', () => {
        expect(partitionLabels('aabbb')).toEqual([2, 3]);
    });

    it('example 2', () => {
        expect(partitionLabels('eccbbbbdec')).toEqual([10]);
    });

    it('interleaved', () => {
        expect(partitionLabels('abac')).toEqual([3, 1]);
    });
});
