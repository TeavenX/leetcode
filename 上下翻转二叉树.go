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
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := root.Right
	root.Right = nil
	left := root.Left
	if right == nil && left == nil {
		// 最后一层
		return root
	}
	root.Left = nil
	next := upsideDownBinaryTree(left)
	left.Left = right
	left.Right = root
	return next
}
