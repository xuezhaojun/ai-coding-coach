# ACM: EOF Multi-Case Input

- **Difficulty**: Easy
- **Category**: ACM Input/Output
- **Topics**: bufio, EOF detection, multiple test cases
- **Link**: [牛客网华为机试](https://www.nowcoder.com/ta/huawei)

## Description

This is an ACM-mode practice problem. Write a complete Go program that reads from `stdin` and writes to `stdout`.

Given multiple test cases, each consisting of two integers `a` and `b`, output the sum of each pair. Read until EOF (end of file).

This is one of the most common ACM input patterns — **process until EOF**.

## Input Format

Multiple lines, each containing two integers `a` and `b`.

```
a1 b1
a2 b2
...
```

There is no line telling you how many cases there are. Read until EOF.

## Output Format

For each line, output `a + b` on its own line.

## Examples

**Example 1:**

```
Input:
1 2
3 4
5 6

Output:
3
7
11
```

## Constraints

- `-10^9 <= a, b <= 10^9`

## Function Signature

Write `solve(reader, writer)` — the core logic.

```go
func solve(reader *bufio.Reader, writer *bufio.Writer)
```

Full ACM wrapper:

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

- **EOF detection**: Use `fmt.Fscan(reader, &a, &b)` and check for `io.EOF`
- This is the single most important ACM pattern for Huawei exams
- Many Huawei problems use this format