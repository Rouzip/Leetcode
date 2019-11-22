// first idea, use heap sort to find n/2 minimum number in array
// wrong, its purpose is not find n/2 elements in the array, it needs to 
// sort the whole array, and select elements interval
func arrayPairSum(nums []int) int {

	length := len(nums)

	for i := length/2 - 1; i >= 0; i-- {
		heapify(nums, length, i)
	}

	res := 0
	for i := length - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	for i := 0; i < length; i++ {
		if i%2 == 0 {
			res += nums[i]
		}
	}
	return res
}

func heapify(arr []int, n int, i int) {
	largest := i
	l, r := 2*i+1, 2*i+2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		// swap two element in array
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}