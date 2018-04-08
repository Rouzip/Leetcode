def searchInsert(nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        if target <= nums[0]:
                return 0
        for i in range(1, len(nums)):
            if nums[i - 1] <= target <= nums[i]:
                return i
        return len(nums)

print(searchInsert([1,3],3))
