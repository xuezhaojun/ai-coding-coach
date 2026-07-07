package string_lines

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read lines until EOF. For each line, reverse the word order.
	// Use bufio.Scanner for line-by-line reading.
	//
	// 适用题型：字符串反转/统计/分割、密码解密、格式化输出、业务模拟
	// 典型输入：若干行字符串，读到 EOF 结束
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