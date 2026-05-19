package sentinel_terminate

import (
	"bufio"
	"fmt"
)

func solveSentinelTerminate(reader *bufio.Reader, writer *bufio.Writer) {
	var a, b int
	for {
		fmt.Fscan(reader, &a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Fprintln(writer, a+b)
	}
}