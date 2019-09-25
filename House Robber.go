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