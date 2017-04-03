def divide(dividend, divisor):
    """
    :type dividend: int
    :type divisor: int
    :rtype: int
    """

    # 判断两个数是否为相同符号
    judge = (dividend > 0) and (divisor > 0)
    result = 0
    # 外层循环保证def divide(self, dividend, divisor):
    """
    :type dividend: int
    :type divisor: int
    :rtype: int
    """

    # 判断两个数是否为相同符号
    judge = (dividend > 0) and (divisor > 0)
    result = 0
    # 外层循环保证除数最终大于被除数
    while dividend >= divisor:
        tmp, i = divisor, 1
        # 内层循环加速减法的进行
        while dividend >= tmp:
            dividend -= tmp
            result += i
            i <<= 1
            tmp <<= 1
    if not judge:
        result = -result
    return min(result, 2147483647) # 除数最终大于被除数


print(divide(-1,1))
