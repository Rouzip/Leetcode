// binary search method
func peakIndexInMountainArray(A []int) int {
	l, r := 0, len(A)-1
	mid := (l+r) / 2

	for l !=r {
		if A[mid-1] < A[mid] && A[mid] > A [mid+1] {
			return mid
		} else if A[mid-1] >= A[mid] {
			r = mid - 1
			mid = (l+r) / 2
		} else {
			l = mid + 1
			mid = (l+r) / 2
		}
	}
	return mid
}

// the most fast method
func peakIndexInMountainArray(A []int) int {
    head, tail := 0, len(A) - 1
    
    for head < tail {
        middle := (head + tail) / 2
        
        if A[middle] < A[middle + 1] {
            head = middle + 1
        } else {
            tail = middle
        }
    }
    
    return tail
}

// use least memory
func peakIndexInMountainArray(A []int) int {
    
    start, end := 0, len(A) - 1
    
    for end - start + 1 > 3 {
        
        mid := start + (end - start) / 2
        
        if A[mid-1] < A[mid] && A[mid] > A[mid+1] {
            
            return mid
        }
        
        if A[mid-1] < A[mid] {
            
            start = mid
        } else {
            
            end = mid
        }
    }
    
    return start + 1
}