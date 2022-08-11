package main

import "fmt"

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	fmt.Println(solveEquation("x=x"))
	fmt.Println(solveEquation("2x=0"))
	fmt.Println(solveEquation("0x=0"))
	fmt.Println(solveEquation("0x=1"))
	fmt.Println(solveEquation("10x=200"))
}

func solveEquation(equation string) string {
	times := 0
	sum := 0
	sign := 1
	num := 0
	addSign := 1
	for i, s := range equation {
		if s == '=' || s == '+' || s == '-' {
			if i > 0 && equation[i-1] != 'x' {
				sum += -1 * addSign * sign * num
				num = 0
			}
			if s == '=' {
				sign = -1
				addSign = 1
			} else if s == '+' {
				addSign = 1
			} else {
				addSign = -1
			}
		} else if s >= '0' && s <= '9' {
			num = 10*num + int(s-'0')
		} else {
			if i == 0 || equation[i-1] > '9' || equation[i-1] < '0' {
				num = 1
			}
			times += addSign * sign * num
			num = 0
		}
	}
	sum += -1 * addSign * sign * num
	//fmt.Println(times, sum)
	if times == 0 && sum == 0 {
		return "Infinite solutions"
	}
	if times == 0 && sum != 0 {
		return "No solution"
	}
	return fmt.Sprintf("x=%d", sum/times)
}
