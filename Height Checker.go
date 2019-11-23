// first idea: sort the whole array and then check the array whether in order or not
func heightChecker(heights []int) int {
	res := 0
	tmp := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		tmp[i] = heights[i]
	}

	sort.Slice(tmp, func(a, b int) bool {
		return tmp[a] < tmp[b]
	})

	for i := 0; i < len(heights); i++ {
		if tmp[i] != heights[i] {
			res++
		}
	}

	return res
}