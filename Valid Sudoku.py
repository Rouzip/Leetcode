def isValidSudoku(board):
    seen = sum(([(c, i), (j, c), (i/3, j/3, c)]
                for i in range(9) for j in range(9)
                for c in [board[i][j]] if c != '.'),[])
    return len(seen) == len(set(seen))


print(isValidSudoku([".87654321","2........","3........","4........","5........","6........","7........","8........","9........"]))
