package matrix_input

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read n rows, m columns, then an n x m matrix.
	// Output the transpose: m rows, n columns, space-separated.
}

// ACM main() template:
// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )
//
// func solve(reader *bufio.Reader, writer *bufio.Writer) {
// 	// your solution here
// }
//
// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()
// 	solve(reader, writer)
// }