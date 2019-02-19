/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {boolean}
 */
var isSameTree = function(p, q) {
  if (p === null && q === null) return true;
  else if (p === null || q === null) return false;
  if (
    p.val === q.val &&
    isSameTree(p.left, q.left) &&
    isSameTree(p.right, q.right)
  )
    return true;
  else return false;
};

/**
 * 在js中判断是否是null需要特别的判断
 */

let isNull = exp => !exp && typeof exp != "undefined" && exp != 0;
