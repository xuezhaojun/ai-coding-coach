package fixed_format

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read n, then n integers, print their sum.
	// Use fmt.Fscan(reader, &x) to read each integer.
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