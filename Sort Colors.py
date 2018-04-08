class Solution:

    def sortColors(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        m = n = 0
        for i in range(len(nums)):
            num = nums[i]
            nums[i] = 2
            if num < 2:
                nums[m] = 1
                m += 1
            if num == 0:
                nums[n] = 0
                n += 1
