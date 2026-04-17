# Merge Triplets to Form Target

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: array, greedy

## Description

Given a list of triplets and a target triplet, determine if you can select triplets and merge them (taking the max of each position) to form the target triplet exactly.

## Approach

Filter out triplets that have any value exceeding the target (they would make the merged result exceed the target). Among remaining triplets, check if each target value can be matched by at least one triplet in that position.
