package graph_input

import "bufio"

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	// Read n nodes, m edges. Then m lines of "u v" edges.
	// Output the out-degree of each node (1 to n), one per line.
	//
	// 适用题型：图的度数/连通性、BFS/DFS遍历、拓扑排序、Dijkstra最短路
	// 典型输入：第一行 n m，接着 m 行每行 u v 表示一条边
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