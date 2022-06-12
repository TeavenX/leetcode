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
func sumRootToLeaf(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var traceback func(node *TreeNode)
	result := 0
	path := 0
	traceback = func(node *TreeNode) {
		path <<= 1
		path |= node.Val
		if node.Left == nil && node.Right == nil {
			result += path
			return
		}
		if node.Left != nil {
			traceback(node.Left)
			path >>= 1
		}
		if node.Right != nil {
			traceback(node.Right)
			path >>= 1
		}
	}
	traceback(root)
	return result
}
