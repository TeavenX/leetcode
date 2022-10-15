package main

import "fmt"

func main() {
	fmt.Println(scoreOfParenthesesV2("(()(()))"))
}

type Stack struct {
	data []int
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Top() int {
	l := s.Len()
	if l > 0 {
		return s.data[l-1]
	}
	return 0
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
	if s.Len() > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return val
	}
	return 0
}

func scoreOfParentheses(s string) int {
	stack := Stack{}
	for _, str := range s {
		if str == '(' {
			stack.Push(0)
		} else {
			cur := stack.Pop()
			pre := stack.Pop()
			if cur == 0 {
				stack.Push(pre + 1)
			} else {
				stack.Push(pre + 2*cur)
			}
		}
	}
	return stack.Top()
}

func scoreOfParenthesesV2(s string) int {
	result := 0
	depth := 0
	for i, str := range s {
		if str == '(' {
			depth++
		} else {
			depth--
			if s[i-1] == '(' {
				result += 1 << depth
			}
		}
	}
	return result
}
