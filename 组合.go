package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	var path []int
	var result [][]int
	// 先声明才能在闭包中递归调用
	var traceback func(int)
	traceback = func(startIdx int) {
		// 判断退出条件
		if len(path) == k {
			// 虽然在go中都是值传递，但是slice、map、channel都是引用类型，传递创建副本时会携带当前指针
			// 所以如果没有创建一个新的变量，前面append的结果会被后面的步骤改变
			// 也可也这么写：
			// temp := make([]int, k)
			// copy(temp, path)
			// result = append(result, temp)
			result = append(result, append([]int{}, path...))
			return
		}
		// 剪枝
		if len(path)+n-startIdx+1 < k {
			return
		}
		// 单层递归逻辑
		for i := startIdx; i <= n; i++ {
			path = append(path, i)
			traceback(i + 1)
			path = path[:len(path)-1]
		}
	}
	traceback(1)
	return result
}
