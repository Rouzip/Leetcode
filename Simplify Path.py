import collections

def simplifyPath(path):
    """
    :type path: str
    :rtype: str
    """
    pathList = path.split('/')
    result = collections.deque()
    for p in pathList:
        if p == '..':
            if result:
                result.pop()
        if p and p != '.':
            result.append(p)
    return '/'+'/'.join(result)


if __name__ == '__main__':
    test = simplifyPath("/..")
    print(test)