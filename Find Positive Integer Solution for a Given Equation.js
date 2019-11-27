/**
 * // This is the CustomFunction's API interface.
 * // You should not implement it, or speculate about its implementation
 * function CustomFunction() {
 *
 *     @param {integer, integer} x, y
 *     @return {integer}
 *     this.f = function(x, y) {
 *         ...
 *     };
 *
 * };
 */
/**
 * @param {CustomFunction} customfunction
 * @param {integer} z
 * @return {integer[][]}
 */
// first idea: brute force loop
// O(n^2)
var findSolution = function(customfunction, z) {
  let res = [];
  for (let i = 0; i <= z; i++) {
    for (let j = 0; j <= z; j++) {
      if (customfunction.f(i, j) === z) {
        res.push([i, j]);
      }
    }
  }
  return res;
};

// the fastest solution
// use two pointer
var findSolution = function(customfunction, z) {
  let x = 1;
  let y = z;
  let res = [];
  while (x > 0 && x <= z && y > 0 && y <= z) {
    if (customfunction.f(x, y) === z) {
      let temp = [x, y];
      res.push(temp);
      x++;
    } else if (customfunction.f(x, y) > z) {
      y--;
    } else {
      x++;
    }
  }
  return res;
};
