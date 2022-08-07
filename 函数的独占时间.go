package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}))
}

type Log struct {
	Ts    int
	Total int
}

type Stack struct {
	data []*Log
}

func (s Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Pop() *Log {
	if len(s.data) == 0 {
		return nil
	}
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}

func (s *Stack) Push(val *Log) {
	s.data = append(s.data, val)
}

func (s *Stack) Top() *Log {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

func exclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	stack := Stack{}
	total := 0
	for _, log := range logs {
		data := strings.Split(log, ":")
		proc, _ := strconv.Atoi(data[0])
		ts, _ := strconv.Atoi(data[2])
		if data[1][0] == 's' {
			logData := &Log{Ts: ts, Total: total}
			stack.Push(logData)
		} else {
			preLog := stack.Pop()
			pop := ts - preLog.Ts + 1 - (total - preLog.Total)
			result[proc] += pop
			total += pop
		}
	}
	return result
}
