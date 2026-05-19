package eof_multi_case

import (
	"bufio"
	"fmt"
	"io"
)

func solveEOFMultiCase(reader *bufio.Reader, writer *bufio.Writer) {
	var a, b int
	for {
		_, err := fmt.Fscan(reader, &a, &b)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		fmt.Fprintln(writer, a+b)
	}
}