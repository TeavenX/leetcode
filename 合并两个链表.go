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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list1
	for i := 0; i < a-1; i++ {
		list1 = list1.Next
	}
	t := list1.Next
	list1.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	for i := a; i <= b; i++ {
		t = t.Next
	}
	list2.Next = t
	return head
}
