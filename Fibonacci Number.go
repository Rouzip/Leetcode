// first idea: dynamic programming
func fib(N int) int {
    if N <= 1 {
        return N
    }
    res := make([]int, N+1)
    res[1] = 1
    for i:=2; i <=N; i++ {
        res[i] = res[i-1] + res[i-2]
    }
    return res[N]
}

// recursion solution, save memory
func fib(N int) int {
	if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

// using Space Optimized Method
func fib(N int) int {
	a, b := 0, 1
	var c, i int
	if N <= 1 {
		return N
	}

	for i := 2; i <= N; i++ {
		c = a + b
		a = b
		b = c
	}
	return b
}