package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

}

func braceExpansionIIError(expression string) []string {
	fmt.Println(expression)
	quote := 0
	left, right := 0, 0
	pre := make([]string, 0)
	for i, b := range expression {
		switch b {
		case '{':
			quote++
			left = i
		case '}':
			quote--
			if quote == 0 {
				right = i
				if len(pre) > 0 {
					if expression[i-1] == ',' {
						pre = union(pre, braceExpansionII(expression[left+1:right]))
					} else {
						pre = traversal(pre, braceExpansionII(expression[left+1:right]))
					}
				} else {
					pre = braceExpansionII(expression[left+1 : right])
				}
			}
		case ',':
		default:
			if quote == 0 {
				pre = append(pre, string(b))
			}
		}
	}
	if left == right {
		pre = strings.Split(expression, ",")
	}
	sort.Strings(pre)
	return pre
}

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
	queue := []string{expression}
	retMap := make(map[string]int, 0)
	for len(queue) != 0 {
		str := queue[0]
		queue = queue[1:]
		// 完成组合
		if strings.Index(str, "{") == -1 {
			retMap[str] = 1
			continue
		}
		// 找到第一组配对的括号
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
