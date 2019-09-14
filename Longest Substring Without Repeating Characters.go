func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	var (
		ans   int
		start int
	)
	s2iMap := make(map[rune]int)
	for end, char := range s {
		pos, contain := s2iMap[char]
		if contain {
			start = max(pos, start)
			delete(s2iMap, char)
		}
		ans = max(ans, end-start+1)

		s2iMap[char] = end + 1
	}
	return ans
}

// 1. 最原始的解法就是不停地回溯循环，找到其中的最优解，由于需要扫描全部的字符
// 每次需要扫描O(j-i)，其次，每次都需要扫到结尾，总共需要扫描n次，因此复杂度为O(n^3)

// 2. 其次就开始使用set，每一次判断是否在该位置的set集合之中可以找到
// 在子串之中检查可以使用set，这样就可以在O(1)的时间内找到是否有子串
// 最终时间复杂度为O(n)

// 3. 使用Map，这样也可以在O(n)时间内找到是否有重复子串