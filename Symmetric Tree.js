/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isSymmetric = function(root) {
  let isSame = (p, q) => {
    if (p === null && q === null) return true;
    else if (p === null || q === null) return false;
    if (p.val === q.val) return true;
    else return false;
  };
  if (root === null) return true;
  let left = [];
  left.push(root.left);
  let right = [];
  right.push(root.right);
  while (left.length && right.length) {
    let node1 = left.pop();
    let node2 = right.pop();
    if (!isSame(node1, node2)) return false;
    left.push(node1.left, node1.right);
    right.push(node2.right, node2.right);
  }
  return true;
};

/**
 * 对于镜像，直接可以想到栈结构的应用。同时参考BFS搜索
 */
