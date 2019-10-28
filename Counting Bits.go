// first idea, map all numbers, change it into binary and count its 1's
// time O(n*sizeof(integer))
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = func(n int) int {
			var res int
			for n != 0 {
				if (n % 2) == 1 {
					res++
				}
				n /= 2
			}
			return res
		}(i)
	}
	return res
}

// combine bit operate and dynamic programming
// time O(n)
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i>>1] + i%2
	}
	return res
}