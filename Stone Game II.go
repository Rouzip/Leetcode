// no idea, cannot find right state transition equation,
// because in my thought, current choice may influence later choice.
// current state has too many choice, i don't know what to choice as the best sub-question solution
// write many question just use one array, i forget to use two dimension array as dp array
func stoneGameII(piles []int) int {
    
}

// dp solution, don't know clearly
// dp[i][j] is the number of stones which Alex can get when starting at index i with M = j
// formula dp[i][j] = max(sum[i]-dp[i+X][max(j,X)], dp[i][j]) where 1<=X<=2*j
func stoneGameII(piles []int) int {
	length := len(piles)
	if length == 0 {
		return 0
	}

	dp := make([][]int, len(piles)+1)
	for i := 0; i <= length; i++ {
		dp[i] = make([]int, length+1)
	}

	sum := make([]int, length+1)

	for i := length - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + piles[i]
	}

	for i := 0; i <= length; i++ {
		dp[i][length] = sum[i]
	}

	for i := length - 1; i >= 0; i-- {
		for j := length - 1; j >= 1; j-- {
			for X := 1; X <= 2*j && i+X <= length; X++ {
				dp[i][j] = max(dp[i][j], sum[i]-dp[i+X][max(j, X)])
			}
		}
	}
	return dp[0][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}