package main

import "fmt"

func main() {
	//root := &TreeNode{Val: 2}
	//root.Left = &TreeNode{Val: 2}
	//root.Right = &TreeNode{Val: 2}
	//root.Left.Left = &TreeNode{Val: 3}
	//root.Right.Left = &TreeNode{Val: 3}
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 0}
	root.Right.Right.Right = &TreeNode{Val: 0}
	fmt.Println(findDuplicateSubtrees(root))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	cache := make(map[string]int)
	result := make([]*TreeNode, 0)
	var dfs func(node *TreeNode) (bool, string)
	dfs = func(node *TreeNode) (bool, string) {
		if node == nil {
			return true, ""
		}
		l, lval := dfs(node.Left)
		r, rval := dfs(node.Right)
		val := fmt.Sprintf("%d(%s)(%s)", node.Val, lval, rval)
		if l && r {
			if cache[val] == 1 {
				result = append(result, node)
				cache[val]++
			} else {
				cache[val]++
			}
		}
		return l && r, val
	}
	dfs(root)
	return result
}

type pair struct {
	node *TreeNode
	idx  int
}

func findDuplicateSubtreesV2(root *TreeNode) []*TreeNode {
	cache := make(map[[3]int]pair)
	result := make([]*TreeNode, 0)
	repeat := map[*TreeNode]struct{}{}
	idx := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		tri := [3]int{node.Val, l, r}
		if p, exist := cache[tri]; exist {
			//result = append(result, p.node)
			repeat[p.node] = struct{}{}
			return p.idx
		}
		idx++
		cache[tri] = pair{node, idx}
		return idx
	}
	dfs(root)
	for node := range repeat {
		result = append(result, node)
	}
	return result
}
