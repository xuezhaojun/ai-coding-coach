# Find Median from Data Stream

- **Difficulty**: Hard
- **Category**: Heap
- **Topics**: two heaps, max-heap, min-heap, design

## Description

Design a data structure that supports adding integers from a data stream and finding the median of all elements added so far. The median is the middle value in a sorted list, or the average of the two middle values if the list has even length.

## Approach

Maintain two heaps: a max-heap for the lower half and a min-heap for the upper half. On each insert, add to the max-heap first, then balance by moving the max of the lower half to the upper half. Keep the max-heap size >= min-heap size. The median is either the max-heap top (odd count) or the average of both tops (even count).
