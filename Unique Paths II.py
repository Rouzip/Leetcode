def uniquePathsWithObstacles(obstacleGrid):
    """
    :type obstacleGrid: List[List[int]]
    :rtype: int
    """
    # m, n = len(obstacleGrid), len(obstacleGrid[0])
    # if not m or not n:
    #     return 0
    # result = [[1 for _ in range(n)] for _ in range(m)]
    # for i in range(1,m):
    #     for j in range(1,n):
    #         result[i][j] = result[i-1][j]+result[i][j-1] if not obstacleGrid[i][j] else 0
    # return result[m-1][n-1]

    m, n = len(obstacleGrid), len(obstacleGrid[0])
    result = [1]+[0]*(n-1)
    for i in range(m):
        for j in range(n):
            if  obstacleGrid[i][j]:
                result[j] = 0
            elif j > 0:
                result[j]+=result[j-1]
    return result[n-1]

if __name__ == '__main__':
    print(uniquePathsWithObstacles([[1],[0]]))
