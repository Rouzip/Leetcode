// first idea: O(m+n)
func intersection(nums1 []int, nums2 []int) []int {
	numMap1 := make(map[int]int)
	numMap2 := make(map[int]int)
	for _, value := range nums1 {
		numMap1[value]++
	}

	for _, value := range nums2 {
		numMap2[value]++
	}

	res := []int{}

	for key := range numMap1 {
		if _, exist := numMap2[key]; exist {
			res = append(res, key)
		}
	}

	return res
}

// least memory usage
// use struct{}{} to reduce memory usage
func intersection(nums1 []int, nums2 []int) []int {
    n1, n2 := nums1, nums2
    
    if len(n1) > len(n2) {
        n1, n2 = n2, n1
    }
    
    m1 := make(map[int]struct{})
    for _, n := range n2 {
        m1[n] = struct{}{}
    }
    
    m2 := make(map[int]struct{})
    for _, n := range n1 {
        if _, ok := m1[n]; ok {
            m2[n] = struct{}{}
        }
    }
    
    res := []int{}
    for n := range m2 {
        res = append(res, n)
    }
    
    return res
}
