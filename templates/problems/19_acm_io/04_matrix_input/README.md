# ACM: Matrix Input

- **Difficulty**: Medium
- **Category**: ACM Input/Output
- **Topics**: bufio, 2D array parsing, matrix operations
- **Link**: [牛客网华为机试](https://www.nowcoder.com/ta/huawei)

## Description

This is an ACM-mode practice problem. Write a complete Go program that reads from `stdin` and writes to `stdout`.

Given an `n x m` matrix, compute the transpose of the matrix. The first line contains two integers `n` (rows) and `m` (columns). The next `n` lines each contain `m` integers representing the matrix.

This pattern (read dimensions then matrix) appears frequently in Huawei exams.

## Input Format

```
n m
a11 a12 ... a1m
a21 a22 ... a2m
...
an1 an2 ... anm
```

## Output Format

The transposed matrix: `m` lines, each with `n` integers separated by spaces.

```
a11 a21 ... an1
a12 a22 ... an2
...
a1m a2m ... anm
```

## Examples

**Example 1:**

```
Input:
2 3
1 2 3
4 5 6

Output:
1 4
2 5
3 6
```

**Example 2:**

```
Input:
3 3
1 0 0
0 1 0
0 0 1

Output:
1 0 0
0 1 0
0 0 1
```

## Constraints

- `1 <= n, m <= 100`
- `-10^9 <= aij <= 10^9`

## Function Signature

Write `solve(reader, writer)` — the core logic.

```go
func solve(reader *bufio.Reader, writer *bufio.Writer)
```

Full ACM wrapper: see 19.02 README for the standard `func main()` template.

## Key Learning Points

- Read `n x m` dimensions first, then read the matrix row by row
- Use a 2D slice `[][]int` for matrix storage
- For output, use `fmt.Fprint(writer, ...)` to build space-separated lines
- **Output format matters**: spaces between numbers, newline at end of each row