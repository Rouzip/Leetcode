// first idea: map the array
// map the arr2, get the order of the second array
// map the arr1, get the order of the first array
func relativeSortArray(arr1 []int, arr2 []int) []int {
	intMap2 := make(map[int]int)
	orderList := []int{}
	
	for i := 0; i < len(arr2); i++ {
		if _, exist := intMap2[arr2[i]]; !exist {
			orderList = append(orderList, arr2[i])
		}
		intMap2[arr2[i]]++
	}

	res := make([]int, len(arr1))
	index := len(arr1) - 1 
	intMap1 := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		intMap1[arr1[i]]++
		if _, exist := intMap2[arr1[i]]; !exist {
			res[index] = arr1[i]
			index--
		}
	}
	
	index = 0
	for i := 0; i < len(orderList); i++ {
		number := orderList[i]
		for j := 0; j < intMap1[number]; j++ {
			res[index] = number
			index++
		}
	}
	sort.Ints(res[index:])

	return res
}