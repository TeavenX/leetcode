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
func evaluateTree(root *TreeNode) bool {
	if root.Val == 0 {
		return false
	} else if root.Val == 1 {
		return true
	}
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == 2 {
		return left || right
	} else {
		return left && right
	}
}
