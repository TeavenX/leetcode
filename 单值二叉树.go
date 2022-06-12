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
func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	var judge func(*TreeNode) bool
	judge = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != val {
			return false
		}
		return judge(node.Left) && judge(node.Right)
	}
	return judge(root)
}
