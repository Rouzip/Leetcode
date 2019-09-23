// math tricks
func divisorGame(N int) bool {
	return N%2 == 0
	// 这里是用到了数学的证明。通过数学归纳法，分类讨论奇数和偶数的情况
}

// dynamic programming
func divisorGame(N int) bool {
	if N <= 1 {
		return false
	}

	dp := make([]bool, N+1)

	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[N]
}