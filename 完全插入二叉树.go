package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 2}
	c := Constructor(root)
	//fmt.Println(c.arr, c.l, c.lazy)
	fmt.Println(c.Insert(2))
	//fmt.Println(c.Insert(3))
	//fmt.Println(c.arr, c.l, c.lazy)
	//fmt.Println(c.Insert(4))
	fmt.Println(c.arr, c.l, c.lazy)
	c.Get_root()
	//fmt.Println(c.arr, c.l, c.lazy)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root *TreeNode
	arr  [4004]int
	l    int
	lazy int
}

func Constructor(root *TreeNode) CBTInserter {
	c := CBTInserter{}
	c.root = root
	c.l = 1
	var dfs func(node *TreeNode, idx int)
	dfs = func(node *TreeNode, idx int) {
		if node == nil {
			return
		}
		c.arr[idx] = node.Val
		c.l++
		dfs(node.Left, idx*2)
		dfs(node.Right, idx*2+1)
	}
	dfs(root, 1)
	c.lazy = c.l
	return c
}

func (this *CBTInserter) Insert(val int) int {
	this.arr[this.l] = val
	p := this.l >> 1
	this.l++
	return this.arr[p]
}

func (this *CBTInserter) Get_root() *TreeNode {
	var dfs func(pre *TreeNode, idx int)
	dfs = func(pre *TreeNode, idx int) {
		fmt.Println("idx:", idx, "lazy:", this.lazy, "this.l:", this.l, "pre:", pre)
		if idx > this.l || pre == nil {
			return
		}
		//if idx&1 == 0 {
		if idx*2 < this.l && idx*2 >= this.lazy {
			if pre.Left == nil {
				pre.Left = &TreeNode{}
			}
			pre.Left.Val = this.arr[idx*2]
		}
		//} else {
		if idx*2+1 < this.l && idx*2+1 >= this.lazy {
			if pre.Right == nil {
				pre.Right = &TreeNode{}
			}
			pre.Right.Val = this.arr[idx*2+1]
		}
		//}
		dfs(pre.Left, idx*2)
		dfs(pre.Right, idx*2+1)
	}
	dfs(this.root, 1)
	this.lazy = this.l
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
