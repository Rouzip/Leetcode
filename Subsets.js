/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function(nums) {
  let result = [
    []
  ];
  nums.forEach((num) => {
    result.push(...result.map(p => p.concat(num)))
  })
  return result;
};

// 思路和python列表推导相同，将每个数组的元素进行遍历，然后将遍历的元素添加到item中去