package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(braceExpansionIICV("{a,b{c,d,{e,f},g}}"))
}

// func braceExpansionIIError(expression string) []string {
// 	fmt.Println(expression)
// 	quote := 0
// 	left, right := 0, 0
// 	pre := make([]string, 0)
// 	for i, b := range expression {
// 		switch b {
// 		case '{':
// 			quote++
// 			left = i
// 		case '}':
// 			quote--
// 			if quote == 0 {
// 				right = i
// 				if len(pre) > 0 {
// 					if expression[i-1] == ',' {
// 						pre = union(pre, braceExpansionII(expression[left+1:right]))
// 					} else {
// 						pre = traversal(pre, braceExpansionII(expression[left+1:right]))
// 					}
// 				} else {
// 					pre = braceExpansionII(expression[left+1 : right])
// 				}
// 			}
// 		case ',':
// 		default:
// 			if quote == 0 {
// 				pre = append(pre, string(b))
// 			}
// 		}
// 	}
// 	if left == right {
// 		pre = strings.Split(expression, ",")
// 	}
// 	sort.Strings(pre)
// 	return pre
// }

func traversal(left, right []string) []string {
	temp := make(map[string]struct{})
	for _, x := range left {
		for _, y := range right {
			temp[x+y] = struct{}{}
		}
	}
	result := make([]string, 0, len(temp))
	for k := range temp {
		result = append(result, k)
	}
	return result
}

func union(left, right []string) []string {
	temp := make(map[string]struct{})
	for _, x := range left {
		temp[x] = struct{}{}
	}
	for _, x := range right {
		temp[x] = struct{}{}
	}
	result := make([]string, 0, len(temp))
	for k := range temp {
		result = append(result, k)
	}
	return result
}

func braceExpansionIICV(expression string) []string {
	// 创建bfs队列，将原始表达式放入队头
	queue := []string{expression}
	// 接收结果的map，这里可以直接写map[string]struct{}
	retMap := make(map[string]int, 0)
	// 开始遍历队列
	for len(queue) != 0 {
		// pop队头
		str := queue[0]
		queue = queue[1:]
		// 如果不存在左/右括号，证明已经是最底层表达式了
		// 将str放到set中
		if strings.Index(str, "{") == -1 {
			retMap[str] = 1
			continue
		}
		// 找到最深一组配对的括号
		left, right := 0, 0
		for i := 0; i < len(str); i++ {
			if str[i] == '{' {
				left = i
			}
			// 找到第一个右括号，与此前最后一个左括号配对
			if str[i] == '}' {
				right = i
				break
			}
		}
		// 取头取尾，注意把当前一对括号去掉
		head := str[:left]
		tail := str[right+1:]
		// 将最大括号里的值按照 , 拆分
		subStrArr := strings.Split(str[left+1:right], ",")
		var strBuilder strings.Builder
		for _, subStr := range subStrArr {
			strBuilder.Reset()
			strBuilder.WriteString(head)
			strBuilder.WriteString(subStr)
			strBuilder.WriteString(tail)
			queue = append(queue, strBuilder.String())
		}
	}
	ret := make([]string, len(retMap))
	i := 0
	for str := range retMap {
		ret[i] = str
		i++
	}
	sort.Strings(ret)
	return ret
}

func braceExpansionII(expression string) []string {
	queue := []string{expression}
	set := make(map[string]struct{}, 0)
	for len(queue) > 0 {
		exp := queue[0]
		queue = queue[1:]
		if strings.Index(exp, "{") == -1 {
			// 没有括号了
			set[exp] = struct{}{}
			continue
		}
		left, right := 0, 0
		for i, b := range exp {
			if b == '{' {
				left = i
			} else if b == '}' {
				right = i
				break
			}
		}
		head := exp[:left]
		tail := exp[right+1:]
		for _, item := range strings.Split(exp[left+1:right], ",") {
			queue = append(queue, head+item+tail)
		}
	}
	ans := make([]string, 0, len(set))
	for k := range set {
		ans = append(ans, k)
	}
	sort.Strings(ans)
	return ans
}
