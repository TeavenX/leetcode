package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	temp := make([]int, 0)
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	var dfs func(arr []int) *TreeNode
	dfs = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		mid := len(arr) >> 1
		root := &TreeNode{Val: arr[mid]}
		root.Left = dfs(arr[:mid])
		root.Right = dfs(arr[mid+1:])
		return root
	}
	return dfs(temp)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBSTV2(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}
	var reInorder func(start, end int) *TreeNode
	reInorder = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := start + (end-start)>>1
		left := reInorder(start, mid-1)
		root := &TreeNode{Val: head.Val, Left: left}
		head = head.Next
		root.Right = reInorder(mid+1, end)
		return root
	}
	return reInorder(0, count-1)
}
