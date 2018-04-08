class Solution(object):

    def dfs(self, nums, path, index, result):
        if len(nums) == index:
            result.append(path)
            return
        for col in range(len(nums)):
            nums[index] = col
            if self.valid(nums, index):
                tmp = '.' * len(nums)
                self.dfs(nums, path + [tmp[:col] + 'Q' +
                                       tmp[col + 1:]], index + 1, result)

    def valid(self, nums, col):
        for i in range(col):
            if abs(nums[col] - nums[i]) == col - i or nums[i] == nums[col]:
                return False
        return True

    def solveNQueens(self, n):
        """
        :type n: int
        :rtype: List[List[str]]
        """

        result = []
        self.dfs([-1] * n, [], 0, result)
        return result


if __name__ == '__main__':
    test = Solution()
    result = test.solveNQueens(4)
    print(result)
