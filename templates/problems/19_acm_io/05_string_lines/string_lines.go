package string_lines

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read lines until EOF. For each line, reverse the word order.
	// Use bufio.Scanner for line-by-line reading.
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