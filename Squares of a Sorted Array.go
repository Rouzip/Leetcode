// combine binary search and binary insert sort
func sortedSquares(A []int) []int {
	length := len(A)
	mid := (length - 1) / 2
	// mid is not correct, not suit for len(A) < 3
	if length == 2 {
		if Abs(A[0]) < Abs(A[1]) {
			return []int {A[0]*A[0], A[1]*A[1]}
		}
		return []int {A[1]*A[1], A[0]*A[0]}
	}
	for (mid-1) > 0 && (mid+1) < len(A)-1 {
		// not correct, not suit for mid in the boundary
		if Abs(A[mid-1]) > Abs(A[mid]) && Abs(A[mid]) < Abs(A[mid+1]) {
			break
		} else {
			if Abs(A[mid-1]) <= Abs(A[mid]) {
				mid = mid - 1
			} else {
				mid = mid + 1
			}
		}
	}

	l, r, index := mid-1, mid+1, 1
	res := make([]int, length)
	res[0] = A[mid] * A[mid]
	for l >= 0 || r < length {
		if l >= 0 && r < length {
			if Abs(A[l]) < Abs(A[r]) {
				res[index] = A[l] * A[l]
				l--
			} else {
				res[index] = A[r] * A[r]
				r++
			}
		} else if r >= length {
			res[index] = A[l] * A[l]
			l--
		} else {
			res[index] = A[r] * A[r]
			r++
		}
		index++
	}

	return res
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// still cannot find right mid
func sortedSquares(A []int) []int {
	length := len(A)
	mid := (length - 1) / 2
	// mid is not correct, not suit for len(A) < 3
	for (mid-1) >= 0 || (mid+1) < length {
		if (mid - 1) >= 0 && (mid + 1) < length {
			if Abs(A[mid-1]) >= Abs(A[mid]) && Abs(A[mid]) <= Abs(A[mid+1]) {
				break
			} else {
				if Abs(A[mid-1]) <= Abs(A[mid]) {
					mid = mid - 1
				} else {
					mid = mid + 1
				}
			}
		} else if (mid - 1) < 0 {
			if Abs(A[mid]) <= Abs(A[mid+1]) {
				break
			}
			mid++
		} else {
			if Abs(A[mid-1]) >= Abs(A[mid]) {
				break
			}
			mid--
		}
	}

	l, r, index := mid-1, mid+1, 1
	res := make([]int, length)
	res[0] = A[mid] * A[mid]
	for l >= 0 || r < length {
		if l >= 0 && r < length {
			if Abs(A[l]) <= Abs(A[r]) {
				res[index] = A[l] * A[l]
				l--
			} else {
				res[index] = A[r] * A[r]
				r++
			}
		} else if r >= length {
			res[index] = A[l] * A[l]
			l--
		} else {
			res[index] = A[r] * A[r]
			r++
		}
		index++
	}

	return res
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// use map the whole array
func sortedSquares(A []int) []int {
	length := len(A)
	mid := 0
	MAX := 1<<31 - 1
	// mid is not correct, not suit for len(A) < 3
	for i := 0; i < length; i++ {
		if Abs(A[i]) < MAX {
			mid = i
			MAX = Abs(A[i])
		}
	}

	index := 1
	res := make([]int, length)
	res[0] = A[mid] * A[mid]
	l, r := mid-1, mid+1
	for l >= 0 || r < length {
		if l >= 0 && r < length {
			if Abs(A[l]) <= Abs(A[r]) {
				res[index] = A[l] * A[l]
				l--
			} else {
				res[index] = A[r] * A[r]
				r++
			}
		} else if r >= length {
			res[index] = A[l] * A[l]
			l--
		} else {
			res[index] = A[r] * A[r]
			r++
		}
		index++
	}

	return res
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// fatest solution
minIndex := 0
minValue := A[0]
for i:=1; i < len(A); i++ {
	compare := A[i]
	if compare < 0 {
		compare *= -1
	}
		
	absMinValue := minValue
	if absMinValue < 0 {
		absMinValue *= -1
	}
	if compare < absMinValue {
		minIndex = i
		minValue = A[i]
	}
}

left := minIndex - 1
right := minIndex + 1

result := []int{A[minIndex]*A[minIndex]}
for len(result) < len(A) {

	// Base case if one goes out of range.
	if left < 0 {
		result = append(result, A[right]*A[right])
		right++
		continue
	} else if right >= len(A) {
		result = append(result, A[left]*A[left])
		left--
		continue
	}
	
	// Handle absolute values
	leftValue := A[left]
	if leftValue < 0 {
		leftValue *= -1
	}
	rightValue := A[right]
	if rightValue < 0 {
		rightValue *= -1
	}
	
	if leftValue <= rightValue {
		result = append(result, A[left]*A[left])
		left--
	} else {
		result = append(result, A[right]*A[right])
		right++
	}
	
}
return result
}