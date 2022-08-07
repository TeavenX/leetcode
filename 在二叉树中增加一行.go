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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	var dfs func(pre, cur *TreeNode, currDepth int)
	dfs = func(pre, cur *TreeNode, currDepth int) {
		if currDepth == depth {
			node := &TreeNode{Val: val}
			if pre == nil {
				node.Left = root
				root = node
			} else {
				if cur == pre.Left {
					node.Left = pre.Left
					pre.Left = node
				} else {
					node.Right = pre.Right
					pre.Right = node
				}
			}
			return
		} else {
			if cur == nil {
				return
			}
			dfs(cur, cur.Left, currDepth+1)
			dfs(cur, cur.Right, currDepth+1)
		}
	}
	dfs(nil, root, 1)
	return root
}
