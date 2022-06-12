package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(permute(a))
}

func permute(nums []int) [][]int {
	n := len(nums)
	path := make([]int, 0)
	result := make([][]int, 0)
	var traceback func()
	traceback = func() {
		if len(path) == n {
			temp := make([]int, n)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < n; i++ {
			used := false
			for j := 0; j < len(path); j++ {
				if path[j] == nums[i] {
					used = true
				}
			}
			if !used {
				path = append(path, nums[i])
				traceback()
				path = path[:len(path)-1]
			}
		}
	}
	traceback()
	return result
}

func permute20220524(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{{n}}
	}
	result := make([][]int, 0)
	path := make([]int, 0)
	var traceback func()
	traceback = func() {
		if len(path) == n {
			result = append(result, append([]int{}, path...))
			return
		}
		for _, num := range nums {
			used := false
			for j := 0; j < len(path); j++ {
				if path[j] == num {
					used = true
				}
			}
			if !used {
				path = append(path, num)
				traceback()
				path = path[:len(path)-1]
			}
		}
	}
	traceback()
	return result
}
