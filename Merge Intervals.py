# Definition for an interval.
class Interval(object):

    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e


class Solution(object):

    def merge(self, intervals):
        """
        :type intervals: List[Interval]
        :rtype: List[Interval]
        """
        if not intervals:
            return intervals
        sortList = sorted(intervals, key=lambda e: e.start)
        result = [sortList[0]]
        tmp = result[0]
        for i in sortList[1:]:
            if tmp.end < i.start:
                result.append(i)
                tmp = result[-1]
            if tmp.end >= i.start:
                tmp.end = max(i.end,tmp.end)
        return result


if __name__ == '__main__':
    test = Solution()
    for i in test.merge([Interval(1, 4), Interval(0, 1)]):
        print(i.start, i.end)
