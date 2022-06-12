package main

import "fmt"

func main() {

	a := constructor([]int{1})
	a = removeNthFromEnd(a, 1)
	a.print()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	slow, fast := pre, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return pre.Next
}

func removeNthFromEnd_1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
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
