# Koko Eating Bananas

- **Difficulty**: Medium
- **Category**: Binary Search
- **Topics**: binary search, math

## Description

Koko loves eating bananas. There are n piles with piles[i] bananas. She has h hours to eat all bananas. Find the minimum integer eating speed k such that she can finish within h hours.

## Approach

Binary search on the eating speed between 1 and the maximum pile size. For each candidate speed, calculate total hours needed. If hours fit within h, try a smaller speed; otherwise try a larger speed.
