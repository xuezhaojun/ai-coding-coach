# N-Queens

- **Difficulty**: Hard
- **Category**: Backtracking
- **Topics**: backtracking, recursion, constraint satisfaction

## Description

Place n queens on an n x n chessboard such that no two queens threaten each other. Return all distinct solutions where each solution is a board configuration represented as a list of strings.

## Approach

Use backtracking row by row. For each row, try placing a queen in each column. Track attacked columns and both diagonals (row-col and row+col) using sets. If a column is safe, place the queen, mark the column and diagonals, and recurse to the next row. Backtrack by removing the queen and unmarking.
