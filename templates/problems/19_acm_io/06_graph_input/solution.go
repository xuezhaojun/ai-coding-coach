package graph_input

import (
	"bufio"
	"fmt"
)

func solveGraphInput(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	degree := make([]int, n+1)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		degree[u]++
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, degree[i])
	}
}