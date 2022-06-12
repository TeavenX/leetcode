package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	path := make([]int, 0)
	n := len(nums)
	var traceback func(int, int)
	traceback = func(startIdx, k int) {
		if k-1 == n {
			return
		}
		for i := startIdx; i < n; i++ {
			path = append(path, nums[i])
			if len(path) == k {
				temp := make([]int, k)
				copy(temp, path)
				result = append(result, temp)
				traceback(i+1, k+1)
				path = path[:len(path)-1]
			}
		}
	}
	traceback(0, 1)
	return result
}
