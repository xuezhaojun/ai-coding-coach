package multiply_strings

// solveMultiply multiplies two numbers represented as strings. O(m*n) time, O(m+n) space.
func solveMultiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	pos := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1, p2 := i+j, i+j+1
			sum := mul + pos[p2]
			pos[p2] = sum % 10
			pos[p1] += sum / 10
		}
	}
	result := make([]byte, 0, m+n)
	for _, v := range pos {
		if len(result) == 0 && v == 0 {
			continue
		}
		result = append(result, byte(v)+'0')
	}
	if len(result) == 0 {
		return "0"
	}
	return string(result)
}
