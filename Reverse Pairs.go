func mergeSortCount(nums []int, begin int, end int) int {
	if begin < end {
		mid := (begin + end) / 2
		count := mergeSortCount(nums, begin, mid) + mergeSortCount(nums, mid+1, end)

		j := mid + 1
		for i := begin; i <= mid; i++ {
			for j <= end && nums[i] > nums[j]*2 {
				j++
			}
			count += j - (mid + 1)
		}
		merge(nums, begin, mid, end)
		return count
	}
	return 0
}

func merge(nums []int, begin int, mid int, end int) {
	n1 := mid - begin + 1
	n2 := end - mid

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = nums[begin+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = nums[mid+i+1]
	}

	i, j := 0, 0

	for m := begin; m <= end; m++ {
		if j >= n2 || (i < n1 && L[i] <= R[j]) {
			nums[m] = L[i]
			i++
		} else {
			nums[m] = R[j]
			j++
		}
	}
}

func reversePairs(nums []int) int {
	return mergeSortCount(nums, 0, len(nums)-1)
}