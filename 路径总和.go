package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	target := targetSum - root.Val
	if target == 0 && root.Left == nil && root.Right == nil {
		return true
	} else if target != 0 {
		return hasPathSum(root.Left, target) || hasPathSum(root.Right, target)
	} else {
		return false
	}
}
