def minPathSum(grid):
    """
    :type grid: List[List[int]]
    :rtype: int
    """
    if not grid:
        return 0    # 如果传入矩阵为空，返回0
    m, n = len(grid), len(grid[0])
    road = [[0 for _ in range(n)] for _ in range(m)]    # 路径选择矩阵
    road[0][0] = grid[0][0]    # 确定开始的长度，防止只有一步
    for i in range(1,m):
        road[i][0] = road[i-1][0]+grid[i][0]
    for i in range(1,n):
        road[0][i] = road[0][i-1]+grid[0][i]
    # 构造路径选择矩阵的第一行与第一列
    for i in range(1,m):
        for j in range(1,n):
            road[i][j] = min(road[i-1][j],road[i][j-1])+grid[i][j]
    return road[-1][-1]

if __name__ == '__main__':
    test = minPathSum([[1,3,1],[1,5,1],[4,2,1]])
    print(test)