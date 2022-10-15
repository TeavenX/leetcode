package main

import "fmt"

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}
	fmt.Println(numComponents(head, []int{4}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
	cache := make(map[int]struct{})
	for _, num := range nums {
		cache[num] = struct{}{}
	}
	result := 0
	pre := false
	for head != nil {
		if _, ok := cache[head.Val]; !ok {
			pre = false
		} else if !pre {
			pre = true
			result++
		}
		head = head.Next
	}
	return result
}
