def combinationSum2(candidates, target):
    """
    :type candidates: List[int]
    :type target: int
    :rtype: List[List[int]]
    """
    sortedList = sorted(candidates)
    result = []
    def dfs(remain, ans, pos):
        if not remain:
            result.append(ans)
            return
        for num, item in enumerate(sortedList):
            if item > remain:
                break
            
            if num < pos or (num > pos and sortedList[num] == sortedList[num - 1]):
                continue
            else:
                dfs(remain - item, ans + [item], num + 1)
    dfs(target, [], 0)
    return result


if __name__ == '__main__':
    a = combinationSum2([10,1,2,7,6,1,5], 8)
    print(a)
