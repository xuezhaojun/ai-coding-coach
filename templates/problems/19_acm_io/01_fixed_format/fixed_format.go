package fixed_format

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read n, then n integers, print their sum.
	// Use fmt.Fscan(reader, &x) to read each integer.
	//
	// 适用题型：数组求和/最大值/平均值、排序题、哈希计数、前缀和
	// 典型输入：第一行 n，第二行 n 个数
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