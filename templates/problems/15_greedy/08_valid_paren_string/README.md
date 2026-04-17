# Valid Parenthesis String

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: string, greedy, dynamic programming

## Description

Given a string containing '(', ')' and '*', where '*' can be treated as '(', ')' or empty string, determine if the string is valid (all parentheses are properly matched).

## Approach

Track the range of possible open parenthesis counts using two variables: lo (minimum) and hi (maximum). For '(' both increase, for ')' both decrease, for '*' lo decreases and hi increases. If hi goes negative, too many closing parens. Clamp lo at 0. Valid if lo is 0 at end.
