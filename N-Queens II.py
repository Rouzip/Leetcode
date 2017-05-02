class Solution(object):

    def dfs(self, nums, index):
        if len(nums) == index:
            self.result += 1
            return
        for col in range(len(nums)):
            nums[index] = col
            if self.valid(nums, index):
                self.dfs(nums, index + 1)

    def valid(self, nums, index):
        for i in range(index):
            if nums[index] == nums[i] or abs(nums[index] - nums[i]) == index - i:
                return False
        return True

    def totalNQueens(self, n):
        """
        :type n: int
        :rtype: int
        """
        self.result = 0

        self.dfs([-1] * n, 0)
        return self.result


if __name__ == '__main__':
    test = Solution()
    test.totalNQueens(4)
    print(test.result)
