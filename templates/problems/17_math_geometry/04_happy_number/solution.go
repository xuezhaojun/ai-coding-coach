package happy_number

// solveIsHappy uses Floyd's cycle detection. O(log n) time, O(1) space.
func solveIsHappy(n int) bool {
	sumDigitSquares := func(num int) int {
		sum := 0
		for num > 0 {
			d := num % 10
			sum += d * d
			num /= 10
		}
		return sum
	}
	slow, fast := n, sumDigitSquares(n)
	for fast != 1 && slow != fast {
		slow = sumDigitSquares(slow)
		fast = sumDigitSquares(sumDigitSquares(fast))
	}
	return fast == 1
}
