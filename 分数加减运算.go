package main

import "fmt"

func main() {
	//fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("-5/2+10/3+7/9"))
}

func fractionAdditionError(expression string) string {
	sign := 1
	numers := make([]int, 0)
	denos := 1
	for idx, exp := range expression {
		if exp == '-' {
			sign = -1
		} else if exp == '+' {
			sign = 1
		} else {
			if exp >= '0' && exp <= '9' {
				num := int(exp - '0')
				if num == 0 && idx > 0 && expression[idx-1] == '1' {
					num = 10
				}
				if idx > 0 && expression[idx-1] == '/' {
					denos *= num
					for i := range numers[:len(numers)-1] {
						numers[i] *= num
					}
				} else {
					fmt.Println(sign, num, denos)
					numers = append(numers, sign*num*denos)
				}
			}
		}
	}
	fmt.Println(numers)
	numer := 0
	result := ""
	for _, num := range numers {
		numer += num
	}
	if numer == 0 {
		return "0/1"
	}
	for i := 2; i <= 7; i++ {
		for denos/i*i == denos && numer/i*i == numer {
			numer /= i
			denos /= i
		}
	}
	result += fmt.Sprintf("%d/%d", numer, denos)
	return result
}

func fractionAddition(expression string) string {
	sign := 1
	numers := make([]int, 0)
	deno := 1
	num := 0
	for idx, exp := range expression {
		switch exp {
		case '-', '+':
			if exp == '-' {
				sign = -1
			} else {
				sign = 1
			}
			if idx > 0 {
				fmt.Println(num)
				deno *= num
				for i := range numers[:len(numers)-1] {
					numers[i] *= num
				}
				num = 0
			}
		case '/':
			fmt.Println(sign, num, deno)
			numers = append(numers, sign*num*deno)
			num = 0
		default:
			num = num*10 + int(exp-'0')
		}
	}
	deno *= num
	for i := range numers[:len(numers)-1] {
		numers[i] *= num
	}
	fmt.Println(numers)
	numer := 0
	result := ""
	for _, num := range numers {
		numer += num
	}
	if numer == 0 {
		return "0/1"
	}
	for i := 2; i <= 7; i++ {
		for deno/i*i == deno && numer/i*i == numer {
			numer /= i
			deno /= i
		}
	}
	result += fmt.Sprintf("%d/%d", numer, deno)
	return result
}
