package sentinel_terminate

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read pairs (a, b). Stop when both are 0.
	// For each pair (except 0 0), print a+b.
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