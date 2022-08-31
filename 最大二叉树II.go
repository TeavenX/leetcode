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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	return helper(node, root, nil)
}

func helper(node, cur, pre *TreeNode) *TreeNode {
	if node.Val > cur.Val {
		node.Left = cur
		if pre != nil {
			pre.Right = node
		}
		return node
	}
	if cur.Right == nil {
		cur.Right = node
	} else {
		cur.Right = helper(node, cur.Right, cur)
	}
	return cur
}

func insertIntoMaxTreeV2(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	var helper func(root *TreeNode) *TreeNode
	helper = func(root *TreeNode) *TreeNode {
		if root == nil {
			return node
		}
		if node.Val > root.Val {
			node.Left = root
			return node
		}
		root.Right = helper(root.Right)
		return root
	}
	return helper(root)
}
