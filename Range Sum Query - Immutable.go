// my own cache idea
type NumArray struct {
	cache []int
}

func Constructor(nums []int) NumArray {
	length := len(nums)
	if length > 0 {
		res := NumArray{cache: make([]int, length)}
		res.cache[0] = nums[0]
		for i := 1; i < length; i++ {
			res.cache[i] = nums[i] + res.cache[i-1]
		}
		return res
	}
	return NumArray{cache: make([]int, 0)}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.cache[j]
	}
	return this.cache[j] - this.cache[i-1]
}

// improved cache idea
// dynamic programming
func Constructor(nums []int) NumArray {
	length := len(nums)
	if length > 0 {
		res := NumArray{cache: make([]int, length+1)}
		for i := 0; i < length; i++ {
			res.cache[i+1] = nums[i] + res.cache[i]
		}
		return res
	}
	return NumArray{cache: make([]int, 0)}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.cache[j+1] - this.cache[i]
}

// we can also use map, the time is O(1)