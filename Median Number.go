// meidian number means that, the number in the left of it is smaller
// than it, and in the right, numbers are bigger than it
func findMid(array []int) (result []bool) {
	length := len(array)

	if length == 1 {
		return []bool{true}
	}
	MIN := -2 ^ 32
	MAX := 2 ^ 32 - 1
	leftMax := make([]int, len(array))
	rightMin := make([]int, len(array))

	leftMax[0] = MIN
	rightMin[length-1] = MAX

	for i := 1; i < length; i++ {
		if leftMax[i-1] > array[i-1] {
			leftMax[i] = leftMax[i-1]
		} else {
			leftMax[i] = array[i-1]
		}
	}
	for i := 1; i < length; i++ {
		if rightMin[length-i] < array[length-i] {
			rightMin[length-i-1] = rightMin[length-i]
		} else {
			rightMin[length-i-1] = array[length-i]
		}
	}

	result = make([]bool, length)
	for i := 0; i < length; i++ {
		if leftMax[i] <= array[i] && rightMin[i] >= array[i] {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}