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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := lists[0]
	lists = lists[1:]
	for len(lists) > 0 {
		head = mergeLists(head, lists[0])
		lists = lists[1:]
	}
	return head
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}
	node := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	}
	if list2 != nil {
		node.Next = list2
	}
	return head
}

func mergeKListsV2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		n := len(lists)
		if n&1 == 1 {
			lists = append(lists, nil)
			n++
		}
		n >>= 1
		for i := 0; i < n; i++ {
			lists[i] = mergeLists(lists[i<<1], lists[i<<1+1])
		}
		lists = lists[:n]
	}
	return lists[0]
}
