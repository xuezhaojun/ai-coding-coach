# ACM: String Lines Processing

- **Difficulty**: Medium
- **Category**: ACM Input/Output
- **Topics**: bufio.Scanner, strings.Fields, line-by-line reading, string parsing
- **Link**: [牛客网华为机试](https://www.nowcoder.com/ta/huawei)

## Description

This is an ACM-mode practice problem. Write a complete Go program that reads from `stdin` and writes to `stdout`.

Given multiple lines of text, each line contains words separated by spaces. Reverse the word order in each line (not the characters within each word, but the order of words), and output the result.

This pattern — reading lines as strings and splitting by whitespace — is extremely common in Huawei exams.

## Input Format

Multiple lines of text, each containing words separated by spaces. Read until EOF.

```
word1 word2 word3 ...
```

## Output Format

For each input line, output the words in reverse order separated by spaces.

```
... word3 word2 word1
```

## Examples

**Example 1:**

```
Input:
hello world
I love coding

Output:
world hello
coding love I
```

**Example 2:**

```
Input:
a b c d
Go is great

Output:
d c b a
great is Go
```

## Constraints

- Each line contains 1-100 words
- Each word contains 1-100 lowercase/uppercase English letters
- Number of lines: 1-100

## Function Signature

Write `solve(reader, writer)` — the core logic.

```go
func solve(reader *bufio.Reader, writer *bufio.Writer)
```

Full ACM wrapper: see 19.02 README for the standard `func main()` template.

Note: This problem uses `bufio.Scanner` inside solve for line-by-line reading, which requires passing `os.Stdin` via reader.

## Key Learning Points

- Use `bufio.Scanner` for line-by-line reading: `scanner.Scan()` reads a line, `scanner.Text()` returns it
- Use `strings.Fields(s)` to split by whitespace (handles multiple spaces)
- Use `strings.Join()` to concatenate with separators
- **`strings.Split(s, " ")` vs `strings.Fields(s)`**: Fields handles multiple consecutive spaces; Split does not
- This is probably the **most common Huawei input pattern**: read lines, split words, process