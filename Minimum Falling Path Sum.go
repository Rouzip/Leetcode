// original idea, just write as the problem define
// don't want to modify input
func minFallingPathSum(A [][]int) int {
	length := len(A)
	dp := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, length+2)
		dp[i][0] = 2<<30 - 1
		dp[i][length+1] = 2<<30 - 1
		if i != 0 {
			for j := 1; j <= length; j++ {
				dp[i][j] = A[i-1][j-1]
			}
		}
	}

	for i := 1; i <= length; i++ {
		for j := 1; j <= length; j++ {
			dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i-1][j+1]) + dp[i][j] 
		}
	}

	res := 2<<30 - 1
	for i := 1; i <= length; i++ {
		res = min(res, dp[length][i])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// the fastest solution
func minFallingPathSum(A [][]int) int {
	N := len(A)
	if N == 0 {
		return 0
	}
	maxInt := 1<<32 - 1
	s := make([][]int, N)
	for i := 0; i < N; i++ {
		s[i] = make([]int, N+2) // including two padding cols
	}
	for i := 0; i < N; i++ {
		s[0][i+1] = A[0][i]
		s[i][0], s[i][N+1] = maxInt, maxInt
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= N; j++ {
			s[i][j] = A[i][j-1] + min(s[i-1][j-1], min(s[i-1][j], s[i-1][j+1]))
		}
	}
	ans := maxInt
	for i := 1; i <= N; i++ {
		ans = min(ans, s[N-1][i])
	}
	return ans
}