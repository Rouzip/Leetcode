def solution(m, n):
	if not m or not n:
		return 0
	result = [[1 for _ in range(n)] for _ in range(m)]
	for i in range(1,m):
		for j in range(1,n):
			result[i][j] = result[i-1][j]+result[i][j-1]
	return result[m-1][n-1]

if __name__ == '__main__':
	print(solution(10,10))