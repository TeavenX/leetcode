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
func widthOfBinaryTreeError(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	maxLen := 1
	emptyNode := 0
	for len(queue) > 0 {
		size := len(queue)
		left, right := 0, size-1
		for left < size && queue[left] == nil {
			left++
		}
		for right >= 0 && queue[right] == nil {
			right--
		}
		emptyNode <<= 1
		cur := right - left + 1 + emptyNode
		if cur > maxLen {
			maxLen = cur
		}
		for i := left; i <= right; i++ {
			if queue[i] != nil {
				queue = append(queue, queue[i].Left, queue[i].Right)
			} else {
				emptyNode++
			}
		}
		queue = queue[size:]
	}
	return maxLen
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type dist struct {
	left, right int
}

func widthOfBinaryTreeError2(root *TreeNode) int {
	cache := map[int]dist{}
	maxLen := 0
	var dfs func(node *TreeNode, level int, idx int)
	dfs = func(node *TreeNode, level int, idx int) {
		if cache[level].left == 0 || cache[level].left > idx {
			p := cache[level]
			p.left = idx
			cache[level] = p
		}
		if cache[level].right < idx {
			p := cache[level]
			p.right = idx
			cache[level] = p
		}
		if node.Left != nil {
			dfs(node.Left, level+1, idx<<1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1, idx<<1|1)
		}
	}
	dfs(root, 0, 1)
	for _, v := range cache {
		l := v.right - v.left + 1
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

func widthOfBinaryTree(root *TreeNode) int {
	cache := map[int]int{}
	maxLen := 0
	var dfs func(node *TreeNode, level int, idx int)
	dfs = func(node *TreeNode, level int, idx int) {
		if cache[level] == 0 {
			cache[level] = idx
		}
		cur := idx - cache[level] + 1
		if cur > maxLen {
			maxLen = cur
		}
		if node.Left != nil {
			dfs(node.Left, level+1, idx<<1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1, idx<<1|1)
		}
	}
	dfs(root, 0, 1)
	return maxLen
}
