package main

import (
	"fmt"
)

func main() {
	str := "[()]{}"
	fmt.Println(isValid(str))
}

type ListNode struct {
	Val  string
	Next *ListNode
}

type Stack struct {
	Top    *ListNode
	Length int
}

func (s *Stack) Push(item string) {
	s.Top = &ListNode{Val: item, Next: s.Top}
	s.Length += 1
}

func (s *Stack) Pop() string {
	if s.Length > 0 {
		result := s.Top.Val
		s.Top = s.Top.Next
		s.Length--
		return result
	}
	return ""
}

func isValid(s string) bool {
	stack := Stack{}
	symbol := map[string]string{"{": "}", "[": "]", "(": ")"}
	for _, str := range s {
		ss := string(str)
		if stack.Length == 0 {
			stack.Push(ss)
		} else {
			if ss == symbol[stack.Top.Val] {
				stack.Pop()
			} else {
				stack.Push(ss)
			}
		}
	}
	return stack.Length == 0
}

func isValidV2(s string) bool {
	stack := make([]byte, 0)
	symbol := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	for _, str := range s {
		strByte := byte(str)
		if len(stack) == 0 {
			stack = append(stack, strByte)
		} else {
			if strByte == symbol[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, strByte)
			}
		}
	}
	return len(stack) == 0
}
