package sentinel_terminate

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read pairs (a, b). Stop when both are 0.
	// For each pair (except 0 0), print a+b.
	//
	// 适用题型：多组查询直到特定条件停止、模拟直到终止信号
	// 典型输入：若干行 a b，遇到 0 0 结束
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