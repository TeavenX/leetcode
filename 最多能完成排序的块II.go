package main

func main() {

}

type Stack struct {
	data []int
}

func (s Stack) Top() int {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1]
	}
	return -1
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
	if len(s.data) > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return val
	}
	return -1
}

func (s Stack) Len() int {
	return len(s.data)
}

func maxChunksToSorted(arr []int) int {
	var head int
	stack := Stack{}
	for _, num := range arr {
		if stack.Len() == 0 || stack.Top() <= num {
			stack.Push(num)
		} else {
			head = stack.Pop()
			for stack.Top() > num {
				stack.Pop()
			}
			stack.Push(head)
		}
	}
	return stack.Len()
}
