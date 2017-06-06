import itertools

def countAndSay(n):
    """
    :type n: int
    :rtype: str
    """
    result = '1'
    for _ in range(n - 1):
        result = ''.join(str(len(list(num))) + digit \
                        for digit, num in itertools.groupby(result))
    return result

print(countAndSay(111))