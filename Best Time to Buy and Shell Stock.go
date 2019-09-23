// brute force to compute the result
func maxProfit(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	length := len(prices)
	resArr := make([]int, length)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			resArr[j] = max(resArr[j], prices[j]-prices[i])
		}
	}
	res := 0
	for i := 0; i < length; i++ {
		res = max(res, resArr[i])
	}
	if res > 0 {
		return res
	}
	return 0
}

// 单顺序查找，通过记录最低值和遍历整个数组，省去两边寻找的时间
// ❌一开始想使用快慢指针，但是有可能两个指针错过，即各自半区的最大最小
// 不一定可以得到全体的最大和最小
func maxProfit(prices []int) int {
	minPri := 2<<30 - 1
	maxProfit := 0
	length := len(prices)

	for i := 0; i < length; i++ {
		if prices[i] < minPri {
			minPri = prices[i]
		}

		if prices[i]-minPri > maxProfit {
			maxProfit = prices[i] - minPri
		}
	}
	if maxProfit < 0 {
		return 0
	}
	return maxProfit
}

// dynamic programming
func maxProfit(prices []int) int {
	length := len(prices)
	res := make([]int, length+1)
}