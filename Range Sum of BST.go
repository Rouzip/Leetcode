/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// recursion method, easy but low effcitive
 func rangeSumBST(root *TreeNode, L int, R int) int {
    if root == nil {
		return 0
	}
	if L <= root.Val && R >= root.Val {
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	} else {
		if L > root.Val {
			return rangeSumBST(root.Right, L, R)
		} else {
			return rangeSumBST(root.Left, L, R)
		}
	}
}

// use tail recursion for better performance
func rangeSumBST(root *TreeNode, L int, R int) int {
    if root == nil{
        return 0
    }
    return helper(root, L, R, 0)
}

func helper(root *TreeNode, L, R, sum int) int{
    if root == nil{
        return sum
    }
    if root.Val >= L && root.Val <= R{
        sum += root.Val
    }
    if root.Val > L{
        sum = helper(root.Left, L, R, sum)
    }
    if root.Val < R{
        sum = helper(root.Right, L, R, sum)
    }
    return sum
}