def spiralOrder(matrix):
    """
    :type matrix: List[List[int]]
    :rtype: List[int]
    """
    result = []
    while matrix:
        result += matrix.pop(0)
        if matrix and matrix[0]:
            for row in matrix:
                result.append(row.pop())
        if matrix:
            result += matrix.pop()[::-1]
        if matrix:
            for row in matrix[::-1]:
                result.append(row.pop(0))
    return result


if __name__ == '__main__':
    test = [[7], [9], [6]]
    print(spiralOrder(test))
