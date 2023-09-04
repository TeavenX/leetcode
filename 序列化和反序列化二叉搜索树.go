package main

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
)

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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	cache := make([]string, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		cache = append(cache, strconv.Itoa(node.Val))
	}
	dfs(root)
	result := strings.Join(cache, ",")
	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	cache := strings.Split(data, ",")
	var contructor func(int, int) *TreeNode
	contructor = func(left, right int) *TreeNode {
		n := len(cache)
		if n == 0 {
			return nil
		}
		rootVal := this.atoi(cache[n-1])
		if rootVal < left || rootVal > right {
			return nil
		}
		cache = cache[:n-1]
		root := &TreeNode{Val: rootVal}
		root.Right = contructor(rootVal, right)
		root.Left = contructor(left, rootVal)
		return root
	}
	root := contructor(math.MinInt, math.MaxInt)
	return root
}

func (this *Codec) atoi(a string) int {
	result, _ := strconv.Atoi(a)
	return result
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	data := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		data = append(data, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	bts, _ := json.Marshal(data)
	return string(bts)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(s string) *TreeNode {
	data := make([]int, 0)
	_ = json.Unmarshal([]byte(s), &data)
	i := 0
	var dfs func(mn, mx int) *TreeNode
	dfs = func(mn, mx int) *TreeNode {
		if i == len(data) {
			return nil
		}
		x := data[i]
		if x < mn || x > mx {
			return nil
		}
		i++
		root := &TreeNode{Val: x}
		root.Left = dfs(mn, x)
		root.Right = dfs(x, mx)
		return root
	}
	return dfs(math.MinInt, math.MaxInt)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
