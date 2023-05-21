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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(node *TreeNode, pre int) bool
	dfs = func(node *TreeNode, pre int) bool {
		if node == nil {
			return false
		}
		pre = pre + node.Val
		if node.Left == nil && node.Right == nil {
			return pre >= limit
		}
		l := dfs(node.Left, pre)
		if !l {
			node.Left = nil
		}
		r := dfs(node.Right, pre)
		if !r {
			node.Right = nil
		}
		return l || r
	}
	r := dfs(root, 0)
	if !r {
		return nil
	}
	return root
}
