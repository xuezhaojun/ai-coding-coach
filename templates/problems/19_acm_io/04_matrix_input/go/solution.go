package matrix_input

import (
	"bufio"
	"fmt"
	"strings"
)

func solveMatrixInput(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m int
	fmt.Fscan(reader, &n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	for j := 0; j < m; j++ {
		parts := make([]string, n)
		for i := 0; i < n; i++ {
			parts[i] = fmt.Sprintf("%d", matrix[i][j])
		}
		fmt.Fprintln(writer, strings.Join(parts, " "))
	}
}