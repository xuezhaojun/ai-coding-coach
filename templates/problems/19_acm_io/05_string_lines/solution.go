package string_lines

import (
	"bufio"
	"fmt"
	"strings"
)

func solveStringLines(reader *bufio.Reader, writer *bufio.Writer) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		words := strings.Fields(line)
		n := len(words)
		for i := n/2 - 1; i >= 0; i-- {
			j := n - 1 - i
			words[i], words[j] = words[j], words[i]
		}
		fmt.Fprintln(writer, strings.Join(words, " "))
	}
}

// Full ACM program:
// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()
// 	solve(reader, writer)
// }