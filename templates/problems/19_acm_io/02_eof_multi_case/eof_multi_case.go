package eof_multi_case

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read pairs (a, b) until EOF. For each pair, print a+b.
	// Use fmt.Fscan in a loop, break on io.EOF.
	//
	// 适用题型：多组数据求和/差/积、多轮模拟、连续查询处理
	// 典型输入：若干行 a b，读到文件末尾为止
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