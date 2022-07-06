package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(evaluate("(let x 2 (mult x (let x 3 y 4 (add x y))))"))
}

const (
	VALUE  = iota // 初始状态
	NONE          // 表达式类型未知
	LET           // let表达式
	LET1          // let已经解析了 vi 变量
	LET2          // let已经解析了最后一个表达式 expr
	ADD           // add表达式
	ADD1          // add表达式已经解析了 e1 表达式
	ADD2          // add表达式已经解析了 e2 表达式
	MULTI         // multi表达式
	MULTI1        // multi表达式已经解析了 e1 表达式
	MULTI2        // multi表达式已经解析了 e2 表达式
	DONE          // 解析完成
)

type expr struct {
	status int
	vr     string // let 的变量 vi
	value  int    // VALUE 的状态数值，或者 LET2 状态最后一个表达式的数值
	e1, e2 int    // add 或者 multi 表达式的两个表达式 e1, e2 的数值
}

func evaluate(expression string) int {
	scope := map[string][]int{} // 作用域
	calculateToken := func(token string) int {
		if unicode.IsLower(rune(token[0])) {
			vals := scope[token]
			return vals[len(vals)-1]
		}
		val, _ := strconv.Atoi(token)
		return val
	}

	vars := make([][]string, 0)
	s := make([]expr, 0)
	cur := expr{status: VALUE}
	for i, n := 0, len(expression); i < n; {
		if expression[i] == ' ' {
			i++ // 跳过空格
			continue
		}
		if expression[i] == '(' {
			i++ // 跳过左括号
			s = append(s, cur)
			// 临时设置状态机为未知表达式
			cur.status = NONE
			continue
		}
		var token string
		if expression[i] == ')' {
			i++ // 跳过右括号
			if cur.status == LET2 {
				token = strconv.Itoa(cur.value)
				for _, vr := range vars[len(vars)-1] { // 清除作用域
					scope[vr] = scope[vr][:len(scope[vr])-1]
				}
				vars = vars[:len(vars)-1]
			} else if cur.status == ADD2 {
				token = strconv.Itoa(cur.e1 + cur.e2)
			} else if cur.status == MULTI2 {
				token = strconv.Itoa(cur.e1 * cur.e2)
			}
			cur, s = s[len(s)-1], s[:len(s)-1] // 弹出上层状态
		} else {
			idxStart := i
			for i < n && expression[i] != ' ' && expression[i] != ')' {
				i++
			}
			token = expression[idxStart:i]
		}
		switch cur.status {
		case VALUE:
			cur.value, _ = strconv.Atoi(token)
			cur.status = DONE
		case NONE:
			switch token {
			case "let":
				cur.status = LET
				vars = append(vars, nil)
			case "add":
				cur.status = ADD
			case "mult":
				cur.status = MULTI
			}
		case LET:
			if expression[i] == ')' {
				// let 表达式最后一个 expr 表达式
				cur.value = calculateToken(token)
				cur.status = LET2
			} else {
				cur.vr = token
				// 记录当前作用域所有变量
				vars[len(vars)-1] = append(vars[len(vars)-1], token)
				cur.status = LET1
			}
		case LET1:
			scope[cur.vr] = append(scope[cur.vr], calculateToken(token))
			cur.status = LET
		case ADD:
			cur.e1 = calculateToken(token)
			cur.status = ADD1
		case ADD1:
			cur.e2 = calculateToken(token)
			cur.status = ADD2
		case MULTI:
			cur.e1 = calculateToken(token)
			cur.status = MULTI1
		case MULTI1:
			cur.e2 = calculateToken(token)
			cur.status = MULTI2
		}
	}
	return cur.value
}
