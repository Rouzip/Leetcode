def lengthOfLastWord(s):
    """
    :type s: str
    :rtype: int
    """
    if not s:
        return 0
    result = s.split(' ')
    for i in range(len(result) - 1, -1, -1):
        if len(result[i]):
            return len(result[i])
    return 0


if __name__ == '__main__':
    print(lengthOfLastWord('a '))
