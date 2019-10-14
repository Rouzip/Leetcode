// O(n), it is a simplify solution of dp
func maxSubArray(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	largest := nums[0]
	cur := largest

	for i := 1; i < length; i++ {
		if nums[i] > nums[i]+cur {
			cur = nums[i]
		} else {
			cur = nums[i] + cur
		}

		if cur > largest {
			largest = cur
		}
	}
	return largest
}

// O(nlogn) division method
func maxSubArray(nums []int) int {
	return findSum(nums, 0, len(nums)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findSum(nums []int, l int, r int) int {
	if l > r {
		return -2 << 30
	}
	m := l + (r-l)/2
	ml, mr := 0, 0

	lMax := findSum(nums, l, m-1)
	rMax := findSum(nums, m+1, r)

	for i, sum := m-1, 0; i >= l; i-- {
		sum += nums[i]
		ml = max(ml, sum)
	}
	for i, sum := m+1, 0; i <= r; i++ {
		sum += nums[i]
		mr = max(mr, sum)
	}

	return max(max(lMax, rMax), ml+mr+nums[m])
}