package fixed_format

import (
	"bufio"
	"fmt"
)

func solveFixedFormat(reader *bufio.Reader, writer *bufio.Writer) {
	var n int
	fmt.Fscan(reader, &n)

	sum := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		sum += x
	}

	fmt.Fprintln(writer, sum)
}