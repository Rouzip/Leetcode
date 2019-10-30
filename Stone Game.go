// first idea, greedy programming, but it cannot handle the situation which maximum number in the middle
func stoneGame(piles []int) bool {
	alex := make([]int, len(piles)/2+1)
	lee := make([]int, len(piles)/2+1)
	stones := piles

	for i := 0; i < len(piles); i++ {
		if (i % 2) == 0 {
			if stones[0] > stones[len(stones)-1] {
				alex[i/2+1] = alex[i/2] + stones[0]
				stones = stones[1:]
			} else {
				alex[i/2+1] = alex[i/2] + stones[len(stones)-1]
				stones = stones[:len(stones)-1]
			}
		} else {
			if stones[0] > stones[len(stones)-1] {
				lee[i/2+1] = lee[i/2] + stones[0]
				stones = stones[1:]
			} else {
				lee[i/2+1] = lee[i/2] + stones[len(stones)-1]
				stones = stones[:len(stones)-1]
			}
		}
	}

	if alex[len(piles)/2] > lee[len(piles)/2] {
		return true
	}
	return false
}

// dynamic programming, each time we can
func stoneGame(piles []int) bool {
	length := len(piles)

	dp := make([][]int, length+2)
	for i := 0; i < length+2; i++ {
		dp[i] = make([]int, length+2)
	}

	for size := 1; size <= length; size++ {
		for i, j := 0, size-1; j < length; i++ {
			parity := (j - i) % 2
			if parity == 1 {
				dp[i+1][j+1] = func(a, b int) int {
					if a > b {
						return a
					}
					return b
				}(piles[i]+dp[i+2][j+1], piles[j]+dp[i+1][j])
			} else {
				dp[i+1][j+1] = func(a, b int) int {
					if a < b {
						return a
					}
					return b
				}(-piles[i]+dp[i+2][j+1], -piles[j]+dp[i+1][j])
			}
			j++
		}
	}

	return dp[1][length] > 0
}

// math method
// we can divide the whole piles into four piles, no matter how lee choose,
// alex always can get bigger 'first + third' or 'second + forth'
func stoneGame(piles []int) bool {
	return true
}