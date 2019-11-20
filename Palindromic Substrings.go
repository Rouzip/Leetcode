// brute force algorithm, compute each part of the whole string
// time complexity O(n^3)
func countSubstrings(s string) int {
	length := len(s)
	res := 0

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if judgeSubString(&s, i, j) {
				res++
			}
		}
	}
	return res
}

func judgeSubString(s *string, start int, end int) bool {
	res := true

	for start <= end {
		if (*s)[start] != (*s)[end] {
			res = false
			break
		}
		start++
		end--
	}
	return res
}

// no idea to use dynamic programming, add a character will make former problem 
// different, how to reuse former problem?

// Solution1
// map each character of the string, since that palindrome has center, so 
// we just need to add result based on last center, we can reduce the time complexity
// to O(n^2)
func countSubstrings(s string) int {
	length, res := len(s), 0

	for center := 0; center <= 2*length-1; center++ {
		left := center / 2
		right := left + center%2
		for left >= 0 && right < length && s[left] == s[right] {
			res++
			left--
			right++
		}
	}
	return res
}

// Solution2
// it is an import solution to judge whether string is palindrome
// Manacher Algorithm
// reuse the result based on palindrome string character
func countSubstrings(s string) int {
	A := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	z := make([]int, len(A))
	center, right := 0, 0

	for i := 1; i < len(z)-1; i++ {
		if i < right {
			z[i] = min(right-i, z[2*center-i])
		}

		for A[i-z[i]-1] == A[i+z[i]+1] {
			z[i]++
		}

		if z[i]+i > right {
			right = z[i] + i
			center = i
		}
	}

	res := 0
	for i := 0; i < len(z); i++ {
		res += (z[i] + 1) / 2
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}