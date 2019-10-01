func flipAndInvertImage(A [][]int) [][]int {
	length := len(A)
	res := make([][]int, length)
	for i := 0; i < length; i++ {
		res[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			res[i][j] = A[i][length-1-j] ^ 1
		}
	}
	return res
}