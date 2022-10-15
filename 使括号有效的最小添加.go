package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("()))(("))
}

type stack struct {
	data []rune
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Push(val rune) {
	s.data = append(s.data, val)
}

func (s *stack) Pop() rune {
	if s.Len() > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return val
	}
	return 0
}

func (s *stack) Top() rune {
	if s.Len() > 0 {
		return s.data[len(s.data)-1]
	}
	return 0
}

func minAddToMakeValid(s string) int {
	st := stack{}
	for _, q := range s {
		if q == '(' {
			st.Push(q)
		} else if st.Top() == '(' {
			st.Pop()
		} else {
			st.Push(q)
		}
	}
	return st.Len()
}

func minAddToMakeValidV2(s string) int {
	flag := 0
	result := 0
	for _, q := range s {
		if q == '(' {
			flag++
		} else {
			if flag > 0 {
				flag--
			} else {
				result++
			}
		}
	}
	return flag + result
}
