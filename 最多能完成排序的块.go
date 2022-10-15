package main

import "fmt"

func main() {
	//arr := []int{1, 2, 0, 3}
	//arr := []int{2, 0, 1}
	arr := []int{2, 1, 0, 3, 4}
	//arr := []int{2, 0, 1, 3, 4}
	fmt.Println(maxChunksToSortedV2(arr))
}

type Stack struct {
	data []int
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(val ...int) {
	s.data = append(s.data, val...)
}

func (s *Stack) Pop() int {
	if s.Len() > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return val
	}
	return -1
}

func (s *Stack) Top() int {
	if s.Len() > 0 {
		return s.data[len(s.data)-1]
	}
	return -1
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func maxChunksToSorted(arr []int) int {
	stack := Stack{}
	for _, num := range arr {
		if stack.IsEmpty() || stack.Top() < num {
			stack.Push(num)
			continue
		}
		max := stack.Top()
		for !stack.IsEmpty() && stack.Top() > num {
			stack.Pop()
		}
		stack.Push(max)
	}
	return stack.Len()
}

func maxChunksToSortedV2(arr []int) int {
	max := 0
	result := 0
	for i, num := range arr {
		if num > max {
			max = num
		}
		if max == i {
			result++
		}
	}
	return result
}
