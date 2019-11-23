// first idea, two pointer to adapt the array
func sortArrayByParityII(A []int) []int {
	slow, quick := 0, 1
	for slow != len(A) {
		if slow % 2 == 0 {
			// FIXME: wrong, forget to plus 1
			if A[slow] % 2 == 0 {
				continue
			} else {
				for A[quick] % 2 != 0 {
					quick++
				}
				A[quick], A[slow] = A[slow], A[quick]
			}
		} else {
			// FIXME: wrong, forget to plus 1
			if A[slow] % 2 == 1 {
				continue
			} else {
				for A[quick] % 2 != 1 {
					quick++
				}
				A[quick], A[slow] = A[slow], A[quick]
			}
		}
		slow++
		// FIXME: wrong, forget to plus 1
		// forget to put quick pointer after slow
	}

	return A 
}

// correct solution
func sortArrayByParityII(A []int) []int {
	slow, quick := 0, 1
	for slow < len(A) {
		if slow%2 == 0 {
			if A[slow]%2 != 0 {
				for A[quick]%2 != 0 {
					quick++
				}
				A[quick], A[slow] = A[slow], A[quick]
			}
		} else {
			if A[slow]%2 != 1 {
				for A[quick]%2 != 1 {
					quick++
				}
				A[quick], A[slow] = A[slow], A[quick]
			}
		}
		slow++
		quick = slow + 1
	}

	return A
}

// the fatest solution
func sortArrayByParityII(A []int) []int {
    out := make([]int,len(A))
    odd := int(1)
    even := int(0)
    for _,v := range A {
        if isEven(v) {
            out[even] = v
            even += 2
        } else {
            out[odd] = v
            odd += 2
        }
    }
    return out
}

func isEven(v int) bool{
    if v % 2 == 0 {
        return true
    }
    return false
}