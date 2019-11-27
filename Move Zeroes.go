// first idea: two pointer, slow and quick to exchange the postion of the elements of the array
func moveZeroes(nums []int) {
	slow, quick := 0, 0
	length := len(nums)

	for slow < length {
		if nums[slow] == 0 {
			for quick < length && nums[quick] == 0 {
				quick++
			}
			if quick < length {
				nums[slow], nums[quick] = nums[quick], nums[slow]
			}
		}
		slow++
		quick = slow
	}
}

// fastest solution
// not exchange the numbers, rewrite the array and add 0 in the tail of the array
func moveZeroes(nums []int)  {
    i, j := 0, 0
    for ; j < len(nums); j++ {
        if nums[j] != 0 { 
            if i != j { nums[i] = nums[j] }
            i++
        }
    }
    for ; i < len(nums); i++ {
        nums[i] = 0
    }
}

// least memory usage solution
func moveZeroes(nums []int)  {
    count := 0
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		} else {
			count++
		}
	}

	//fmt.Printf("nums: %v count: %v, index: %d\n", nums, count, index)
	for i := 0; i < count; i++ {
		nums[index] = 0
		index++
    }   
}