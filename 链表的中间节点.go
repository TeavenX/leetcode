package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := constructor(a)
	ret := middleNodeDoublePtr(b)
	ret.print()
}

func constructor(s []int) *ListNode {
	head := ListNode{Val: s[0]}
	ptr := &head
	for _, item := range s[1:] {
		temp := ListNode{Val: item}
		ptr.Next = &temp
		ptr = &temp
	}
	return &head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	temp := l
	fmt.Printf("%d", temp.Val)
	temp = temp.Next
	for temp != nil {
		fmt.Printf("->%d", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func middleNode(head *ListNode) *ListNode {
	i := 0
	ptr := head
	for ptr != nil {
		i += 1
		ptr = ptr.Next
	}
	i = i / 2
	for i > 0 {
		head = head.Next
		i--
	}
	return head
}
func middleNodeDoublePtr(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode20220629(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	return slow
}
