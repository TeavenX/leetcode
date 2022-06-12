package main

import "fmt"

func main() {
	fmt.Println(canIWin(4, 6))
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= maxChoosableInteger {
		return true
	}
	if ((1+maxChoosableInteger)*maxChoosableInteger)>>1 < desiredTotal {
		return false
	}
	cache := make(map[int]bool)
	var dfs func(int, int) bool
	dfs = func(used, cur int) bool {
		if exist := cache[used]; !exist {
			result := false
			for i := 0; i < maxChoosableInteger; i++ {
				if (used>>i)&1 == 0 {
					// 当前数字没有被选择
					if cur+i+1 >= desiredTotal {
						result = true
						break
					}
					if !dfs(used|1<<i, cur+i+1) {
						result = true
						break
					}
				}
			}
			cache[used] = result
		}
		return cache[used]
	}
	return dfs(0, 0)
}
