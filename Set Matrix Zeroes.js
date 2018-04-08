/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
var setZeroes = function(matrix) {
  let rowList = new Set();
  let columnList = new Set();
  matrix.forEach((row, rowLoc) => {
    row.forEach((column, columnLoc) => {
      if (column == 0) {
        rowList.add(rowLoc);
        columnList.add(columnLoc);
      }
    })
  })
  if (!rowList.size)
    return;
  console.log(rowList, columnList);
  matrix.forEach((row, rowLoc) => {
    row.forEach((column, columnLoc) => {
      if (rowList.has(rowLoc) || columnList.has(columnLoc))
        matrix[rowLoc][columnLoc] = 0;
    })
  })
};

// 记录了为0的行和列，然后统一循环置0
a = [
  [0, 0, 0, 5],
  [4, 3, 1, 4],
  [0, 1, 1, 4],
  [1, 2, 1, 3],
  [0, 0, 1, 1]
]
setZeroes(a)
console.log(a);