func balancedStringSplit(s string) int {
	left := 0
	right := 0
	result := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			left++
		} else if s[i] == 'R' {
			right++
		}
		if left == right && left != 0 {
			result++
		}
	}

	return result
}