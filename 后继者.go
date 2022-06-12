package main

import "fmt"

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	fmt.Println(inorderSuccessor(root, root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	var pre *TreeNode
	var target *TreeNode
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre == p {
			target = root
		}
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	return target
}
