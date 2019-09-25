// dynamic programme but not elegant
func rob(nums []int) int {
	length := len(nums)
	if length <= 3 {
		if length == 0 {
			return 0
		}
		if length == 1 {
			return nums[0]
		} else if length == 2 {
			if nums[0] > nums[1] {
				return nums[0]
			}
			return nums[1]
		} else {
			if nums[0]+nums[2] > nums[1] {
				return nums[0] + nums[2]
			}
			return nums[1]
		}
	}

	res := make([]int, len(nums))
	res[0], res[1], res[2] = nums[0], nums[1], nums[0]+nums[2]

	for i := 3; i < length; i++ {
		if res[i-3] > res[i-2] {
			res[i] = res[i-3] + nums[i]
		} else {
			res[i] = res[i-2] + nums[i]
		}
	}

	if res[length-2] > res[length-1] {
		return res[length-2]
	}
	return res[length-1]
}

// a elegant solution
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	res := make([]int, 1+length)
	res[0], res[1] = 0, nums[0]

	for i := 2; i <= length; i++ {
		if res[i-1] > (res[i-2] + nums[i-1]) {
			res[i] = res[i-1]
		} else {
			res[i] = res[i-2] + nums[i-1]
		}
	}

	return res[length]
}

// reduce space
func rob(nums []int) int {
	fir, sec := 0, 0

	for i := len(nums) - 1; i >= 0; i-- {
		var sec_next int
		fir_next := nums[i] + sec
		if fir > sec {
			sec_next = fir
		} else {
			sec_next = sec
		}
		fir = fir_next
		sec = sec_next
	}
	if fir > sec {
		return fir
	}
	return sec
}