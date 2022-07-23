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
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(node *TreeNode) (bool, *TreeNode)
	dfs = func(node *TreeNode) (bool, *TreeNode) {
		if node == nil {
			return false, nil
		}
		leftContainOne, left := dfs(node.Left)
		rightContainOne, right := dfs(node.Right)
		if node.Val == 0 && !leftContainOne && !rightContainOne {
			return false, nil
		}
		node.Left, node.Right = left, right
		return true, node
	}
	_, root = dfs(root)
	return root
}
