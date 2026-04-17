# Valid Sudoku

- **Difficulty**: Medium
- **Category**: Arrays & Hashing
- **Topics**: array, hash table, matrix

## Description

Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the rules: each row, column, and 3x3 sub-box must contain the digits 1-9 without repetition.

## Approach

Use three arrays of boolean sets (one for rows, columns, and 3x3 boxes). Iterate through every cell; for each digit found, check if it has already appeared in its row, column, or box. The box index is computed as (row/3)*3 + col/3.
