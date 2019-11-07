// first idea, O(n^2)
func minDeletionSize(A []string) int {
	result := 0
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}
	for i := 0; i < len(A[0]); i++ {
		for j := 1; j < len(A); j++ {
			if A[j][i] < A[j-1][i] {
				result++
				break
			}
		}
	}
	return result
}

