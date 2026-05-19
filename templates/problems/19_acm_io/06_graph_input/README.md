# ACM: Graph Input (Adjacency List)

- **Difficulty**: Medium
- **Category**: ACM Input/Output
- **Topics**: bufio, adjacency list construction, graph parsing
- **Link**: [牛客网华为机试](https://www.nowcoder.com/ta/huawei)

## Description

This is an ACM-mode practice problem. Write a complete Go program that reads from `stdin` and writes to `stdout`.

Given a directed graph with `n` nodes (numbered 1 to n) and `m` edges, count the out-degree of each node. The first line contains two integers `n` and `m`. The next `m` lines each contain two integers `u` and `v` representing a directed edge from `u` to `v`.

This pattern (read n/m then edge list) is used in virtually every Huawei graph problem.

## Input Format

```
n m
u1 v1
u2 v2
...
um vm
```

## Output Format

`n` lines, each containing the out-degree of node `i` (for i = 1 to n).

```
degree_of_1
degree_of_2
...
degree_of_n
```

## Examples

**Example 1:**

```
Input:
4 5
1 2
1 3
2 3
2 4
3 4

Output:
2
2
1
0
```

**Example 2:**

```
Input:
3 2
1 2
3 2

Output:
1
0
1
```

## Constraints

- `2 <= n <= 10^5`
- `1 <= m <= 10^5`
- `1 <= u, v <= n`

## Function Signature

Write `solve(reader, writer)` — the core logic.

```go
func solve(reader *bufio.Reader, writer *bufio.Writer)
```

Full ACM wrapper: see 19.02 README for the standard `func main()` template.

## Key Learning Points

- **Graph input format**: First line `n m`, then `m` lines of edges — this is the standard format in Chinese competitive programming
- Build adjacency list as `[][]int` (slice of slices)
- Node numbering from 1 to n means allocate `n+1` sized slice
- This pattern directly maps to Course Schedule, Number of Islands (with edge input), Dijkstra, etc.
- **In-degree vs out-degree**: Keep separate arrays when doing topological sort