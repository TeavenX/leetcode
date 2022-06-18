package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: -3}
	//root.Right = &TreeNode{Val: -5}
	fmt.Println(findFrequentTreeSum(root))
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
func findFrequentTreeSum(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	cache := make(map[int]int)
	maxCount := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := dfs(root.Left) + root.Val + dfs(root.Right)
		cache[sum]++
		if cache[sum] > maxCount {
			maxCount = cache[sum]
		}
		return sum
	}
	dfs(root)
	result := make([]int, 0)
	for k, v := range cache {
		if v == maxCount {
			result = append(result, k)
		}
	}
	return result
}
