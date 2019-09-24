func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	costArr := make([]int, length+1)
	if len(cost) < 2 {
		return cost[0]
	}

	costArr[0], costArr[1] = cost[0], cost[1]

	for i := 2; i < length; i++ {
		if costArr[i-1] < costArr[i-1] {
			costArr[i] = costArr[i-1] + cost[i]
		} else {
			costArr[i] = costArr[i-2] + cost[i]
		}
	}
	if costArr[length-2] < costArr[length-1] {
		return costArr[length-2]
	}
	return costArr[length-1]
}