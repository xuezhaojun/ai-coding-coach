# ACM: Fixed Format Input

- **Difficulty**: Easy
- **Category**: ACM Input/Output
- **Topics**: bufio, fmt.Fscan, fixed format parsing
- **Link**: [牛客网华为机试](https://www.nowcoder.com/ta/huawei)

## Description

This is an ACM-mode practice problem. You must write a complete Go program that reads from `stdin` and writes to `stdout`.

Given a list of integers, calculate the sum of all elements. The first line contains an integer `n` — the number of elements. The second line contains `n` integers.

## Input Format

```
n
a1 a2 a3 ... an
```

## Output Format

A single integer: the sum of all elements.

## Examples

**Example 1:**

```
Input:
5
1 2 3 4 5

Output:
15
```

**Example 2:**

```
Input:
3
-1 0 1

Output:
0
```

## Constraints

- `1 <= n <= 10^5`
- `-10^9 <= ai <= 10^9`

## Function Signature

Write `solve(reader, writer)` — the core logic extracted from a full ACM program.

```go
func solve(reader *bufio.Reader, writer *bufio.Writer)
```

When running on a real ACM platform, wrap it in `func main()`:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// your solution here
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	solve(reader, writer)
}
```

## Key Learning Points

- Use `bufio.NewReader(os.Stdin)` for fast input, always use bufio to avoid TLE
- Use `fmt.Fscan(reader, &n)` to read integers
- Use `bufio.NewWriter(os.Stdout)` + `defer writer.Flush()` for output
- The `solve()` pattern lets you test locally with `strings.NewReader`