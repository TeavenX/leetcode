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
func nextLargerNodes(head *ListNode) []int {
	stack := make([]int, 0)
	var ans []int
	var dfs func(node *ListNode, i int)
	dfs = func(node *ListNode, i int) {
		if node == nil {
			ans = make([]int, i)
			return
		}
		dfs(node.Next, i+1)
		for len(stack) > 0 && stack[len(stack)-1] <= node.Val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, node.Val)
	}
	dfs(head, 0)
	return ans
}

func nextLargerNodesV2(head *ListNode) []int {
	type pair struct {
		val, i int
	}
	stack := make([]pair, 0)
	ans := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		v := node.Val
		for len(stack) > 0 && stack[len(stack)-1].val < v {
			// 当栈顶的值小于当前node的值时，意味着当前节点已经构成了前面至少一个值的下一个更大值条件了
			// 此时循环将小于当前node值的节点弹出，设置前面节点的更大值为当前node值
			// 由于要设置前面节点的值，栈又是单调的，因此栈内需要记录前面值的索引（idx）
			ans[stack[len(stack)-1].i] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, len(ans)})
		ans = append(ans, 0)
	}
	return ans
}
