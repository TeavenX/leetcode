package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(myAtoi("  -+123"))
	fmt.Println(myAtoiFSM("9223372036854775808"))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := int64(0)
	flag := false
	sign := 1
	for _, str := range s {
		if str == ' ' {
			if flag {
				break
			}
			continue
		}
		num := str - '0'
		if !flag && (str == '+' || str == '-') {
			flag = true
			if str == '-' {
				sign = -1
			}
		} else if num <= 9 && num >= 0 {
			flag = true
			result = result*10 + int64(num)
			if result > math.MaxInt32 {
				break
			}
		} else {
			break
		}
	}
	if result > math.MaxInt32 || result < 0 {
		if sign == 1 {
			result = math.MaxInt32
		} else {
			result = -(math.MaxInt32 + 1)
		}
	} else {
		result = int64(sign) * result
	}
	return int(result)
}

const (
	START_STATE  = "start"
	SIGN_STATE   = "sign"
	NUMBER_STATE = "number"
	END_STATE    = "end"
)

type State interface {
	Enter(b rune)
	Transfer(b rune) string
}

type StartState struct{}

func (s *StartState) Enter(b rune) {}
func (s *StartState) Transfer(b rune) string {
	switch b {
	case '+', '-':
		return SIGN_STATE
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER_STATE
	case ' ':
		return START_STATE
	default:
		return END_STATE
	}
}

type SignState struct {
	sign int
}

func (s *SignState) Enter(b rune) {
	if b == '+' {
		s.sign = 1
	} else {
		s.sign = -1
	}
}
func (s *SignState) Transfer(b rune) string {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER_STATE
	default:
		return END_STATE
	}
}

type NumberState struct {
	num int
}

func (s *NumberState) Enter(b rune) {
	s.num *= 10
	s.num += int(b - '0')
}
func (s *NumberState) Transfer(b rune) string {
	if s.num > math.MaxInt32 {
		return END_STATE
	}
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER_STATE
	default:
		return END_STATE
	}
}

type EndState struct{}

func (s *EndState) Enter(b rune)           {}
func (s *EndState) Transfer(b rune) string { return "end" }

func myAtoiFSM(s string) int {
	mp := map[string]State{
		START_STATE:  &StartState{},
		SIGN_STATE:   &SignState{sign: 1},
		NUMBER_STATE: &NumberState{},
		END_STATE:    &EndState{},
	}
	state := mp[START_STATE]
	for _, b := range s {
		next := state.Transfer(b)
		if next == END_STATE {
			break
		}
		state = mp[next]
		state.Enter(b)
	}
	ans := interface{}(mp[SIGN_STATE]).(*SignState).sign * interface{}(mp[NUMBER_STATE]).(*NumberState).num
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	} else if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}
