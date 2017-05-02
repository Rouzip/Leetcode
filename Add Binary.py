def addBinary(a, b):
    """
    :type a: str
    :type b: str
    :rtype: str
    """
    if not a:
        return b
    elif not b:
        return a

    if a[-1] == '1' and b[-1] == '1':
        return addBinary(addBinary(a[0:-1], b[0:-1]), '1') + '0'
    elif a[-1] == '0' and b[-1] == '0':
        return addBinary(a[0:-1], b[0:-1]) + '0'
    else:
        return addBinary(a[0:-1], b[0:-1]) + '1'


if __name__ == '__main__':
    print(addBinary('1010', '1011'))
