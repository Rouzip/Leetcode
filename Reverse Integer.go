func reverse(x int) int {
	var (
		result int64
		mod    int64
	)
	tail := int64(x)
	for {
		mod = tail % 10
		tail = tail / 10
		result = 10*result + mod
		if tail == 0 {
			if result > (2<<30-1) || result < -(2<<30) {
				return 0
			}
			return int(result)
		}
	}
}

// 思路1. 转换为字符串，之后转置字符串，去掉0
// 思路2. 除以10，再加
// 注意的地方：因为是补码，所以在去掉一位符号位的时候，还需要思考到一半为正，一半为负