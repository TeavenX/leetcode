package main

import "math/rand"

func main() {

}

type Node struct {
	Val  int
	Next *Node
	Down *Node
}

type Skiplist struct {
	Head *Node
}

func Constructor() Skiplist {
	return Skiplist{Head: &Node{Val: -1}}
}

func (this *Skiplist) Search(target int) bool {
	//先在顶层查找
	cur := this.Head
	for cur != nil {
		// cur比目标值大时，进入下一层没有意义，必须要cur的下一个节点比目标值大，然后在cur进入下一层
		for cur.Next != nil && cur.Next.Val < target {
			cur = cur.Next
		}
		// 如果找到直接返回，如果目标值比当前层的最大值还大也直接返回（不可能出现在下一层）
		if cur.Next != nil && cur.Next.Val == target {
			return true
		}
		// 否则进入下一层
		cur = cur.Down
	}
	return false
}

func (this *Skiplist) Add(num int) {
	cur, isInsert := this.Head, true
	down := &Node{Val: -1}
	// 设置下沉队列
	queue := make([]*Node, 0)
	for cur != nil {
		// 跟search相同的思路
		for cur.Next != nil && cur.Next.Val < num {
			cur = cur.Next
		}
		// 当前层的前置节点入队列
		queue = append(queue, cur)
		// 进入下一层，一直进到底层
		cur = cur.Down
	}
	// isInsert标记位设置是否更新跳表的上层索引
	// 判断是否有定位到插入节点的上一个节点
	// 如果有定位到，queue的最后一个元素就是当前层的待插入节点的上一个节点，
	for isInsert && len(queue) > 0 {
		// 从queue中弹出待修改后继的下一个节点
		cur = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if down.Val == -1 {
			// down的值是-1，证明当前处于跳表的最底层，只需要修改后继，下一层置为nil
			cur.Next = &Node{Val: num, Next: cur.Next, Down: nil}
		} else {
			// down有值，证明是在更新上层索引的阶段
			cur.Next = &Node{Val: num, Next: cur.Next, Down: down}
		}
		// 把down修改为num对应的节点，提交给上层索引
		down = cur.Next
		// 设置更新上层索引概率为50%
		isInsert = rand.Float64() < 0.5
	}
	if isInsert {
		// 50%的概率将头节点拔高一层
		this.Head = &Node{Val: -1, Down: this.Head}
	}
}

func (this *Skiplist) Erase(num int) bool {
	cur, isFound := this.Head, false
	for cur != nil {
		for cur.Next != nil && cur.Next.Val < num {
			cur = cur.Next
		}
		if cur.Next != nil && cur.Next.Val == num {
			isFound = true
			cur.Next = cur.Next.Next
		}
		cur = cur.Down
	}
	return isFound
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
