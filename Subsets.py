class Solution:

    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        from copy import deepcopy
        result = []

        def getResult(curr, end, tmpResult):
            next_loc = curr + 1
            if curr == end:
                result.append(deepcopy(tmpResult))
            else:
                getResult(next_loc, end, tmpResult)
                tmp = deepcopy(tmpResult)
                tmp.append(nums[curr])
                getResult(next_loc, end, tmp)
        getResult(0, len(nums), [])
        return result

    def subsets2(self, nums):
        result = [[]]
        for num in nums:
            result.extend([item + [num] for item in result])
        return result


if __name__ == '__main__':
    a = Solution()
    print(a.subsets3([1, 2, 3]))
