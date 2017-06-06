def searchMatrix(matrix, target):
    """
    :type matrix: List[List[int]]
    :type target: int
    :rtype: bool
    """
    if not matrix or target == None:
        return False
    col,row = len(matrix[0]),len(matrix)
    high,low = col*row-1,0
    while low<=high:
        mod = int((high+low)/2)
        if matrix[int(mod/col)][mod%col] == target:
            return True
        elif matrix[int(mod/col)][mod%col] < target:
            low = mod+1
        else:
            high = mod-1
    return False
        

if __name__ == '__main__':
    test = searchMatrix([],0)
    print(test)
