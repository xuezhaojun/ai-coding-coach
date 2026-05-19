# ACM: Sentinel Termination Input

- **Difficulty**: Easy
- **Category**: ACM Input/Output
- **Topics**: bufio, sentinel value, conditional loop termination
- **Link**: [牛客网华为机试](https://www.nowcoder.com/ta/huawei)

## Description

This is an ACM-mode practice problem. Write a complete Go program that reads from `stdin` and writes to `stdout`.

Given multiple pairs of integers `a` and `b`, output the sum of each pair. Stop reading when you encounter `0 0` (this pair should NOT be processed).

This is the second most common ACM pattern — **read until a sentinel value**.

## Input Format

Multiple lines, each containing two integers. The last line is `0 0`.

```
a1 b1
a2 b2
...
0 0
```

## Output Format

For each line (excluding `0 0`), output `a + b` on its own line.

## Examples

**Example 1:**

```
Input:
1 2
3 4
0 0

Output:
3
7
```

**Example 2:**

```
Input:
0 0

Output:

(no output)
```

## Constraints

- `-10^9 <= a, b <= 10^9`
- Input is guaranteed to end with `0 0`

## Function Signature

Write `solve(reader, writer)` — the core logic.

```go
func solve(reader *bufio.Reader, writer *bufio.Writer)
```

Full ACM wrapper: see 19.02 README for the standard `func main()` template.

## Key Learning Points

- **Sentinel detection**: Read and check before processing, break on match
- Different from EOF: you check for a specific condition, not end-of-file
- Common variants: terminate on `0`, on negative number, on specific string