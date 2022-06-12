package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	for list1 != nil {
		cur.Next = list1
		list1 = list1.Next
		cur = cur.Next
	}
	for list2 != nil {
		cur.Next = list2
		list2 = list2.Next
		cur = cur.Next
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists20220529(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	result := &ListNode{}
	temp := result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			temp.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		temp = temp.Next
	}
	for list1 != nil {
		temp.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		temp = temp.Next
	}
	for list2 != nil {
		temp.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		temp = temp.Next
	}
	return result.Next
}
