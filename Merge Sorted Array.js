/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
  nums1 = [...nums1.slice(0, m), ...nums2.slice(0, n)]
    .sort((a, b) => a - b)
    .forEach((value, index) => (nums1[index] = value));
};

/**
 * 函数排序比较的是ASCII码，所以需要传入比较函数
 * js函数为值传递
 * 数组与对象传入的是引用，如果直接改变引用，改变的是局部变量
 */
