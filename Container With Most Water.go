// brute force solution O(n^2)
func maxArea(height []int) int {
	length := len(height)

	var answer = -2 << 30
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			min := func(a, b int) int {
				if a > b {
					return b * b
				}
				return a * a
			}(height[i], height[j])
			tmp := min * min * (j - i)
			if tmp > answer {
				answer = tmp
			}
		}
	}
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// two side map O(n)
// we need to discuss different situations

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	var answer int

	for l < r {
		answer = max(answer, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return answer
}