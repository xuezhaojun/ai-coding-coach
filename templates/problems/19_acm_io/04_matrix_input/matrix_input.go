package matrix_input

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read n rows, m columns, then an n x m matrix.
	// Output the transpose: m rows, n columns, space-separated.
	//
	// 适用题型：矩阵转置/旋转、二维前缀和、网格BFS/DFS、图像处理
	// 典型输入：第一行 n m，接着 n 行每行 m 个数
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