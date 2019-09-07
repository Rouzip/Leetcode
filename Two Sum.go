func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	result := make([]int, 2)
	for index, v := range nums {
		_, exist := m[v]
		if exist {
			result[0], result[1] = m[v], index
		} else {
			m[target-v] = index
		}
	}
	return result
}

// 利用了hash表来减少循环次数