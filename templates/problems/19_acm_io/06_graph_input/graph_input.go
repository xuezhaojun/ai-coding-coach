package graph_input

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read n nodes, m edges. Then m lines of "u v" edges.
	// Output the out-degree of each node (1 to n), one per line.
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