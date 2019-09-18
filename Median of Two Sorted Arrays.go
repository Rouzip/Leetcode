func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 > len2 {
		temp := nums1
		nums1 = nums2
		nums2 = temp
		tmp := len1
		len1 = len2
		len2 = tmp
	}

	iMin := 0
	iMax := len1
	pos := (len1 + len2 + 1) / 2

	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := pos - i

		if i < iMax && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = Max(nums1[i-1], nums2[j-1])
			}
			if (len1+len2)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == len1 {
				minRight = nums2[j]
			} else if j == len2 {
				minRight = nums1[i]
			} else {
				minRight = Min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}

	return 0.0
}