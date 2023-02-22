package main

func main() {

}

func circularPermutationError(n int, start int) []int {
	mx := 1<<n - 1
	cache := make([]bool, mx+1)
	ans := make([]int, mx+1)
	ans[0] = start
	cache[start] = true
	var traceback func(idx int) bool
	traceback = func(idx int) bool {
		if idx > mx {
			d := abs(ans[idx-1] - start)
			for i := 0; i < n; i++ {
				if d == 1<<i {
					return true
				}
			}
			return false
		}
		pre := ans[idx-1]
		st, ed := 0, n
		if pre&1 == 1 {
			st = 1
			for i := 0; i < n; i++ {
				if 1<<i-1 > pre {
					ed = i - 1
					break
				}
			}
		}
		for i := st; i < ed; i++ {
			flag := -1
			for j := 0; j < 2; j++ {
				flag *= -1
				num := pre + (1<<i)*flag
				if num < 0 || num > mx || cache[num] {
					continue
				}
				cache[num] = true
				ans[idx] = num
				r := traceback(idx + 1)
				if r {
					return true
				}
				cache[num] = false
			}
		}
		return false
	}
	traceback(1)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 格雷码 https://oi-wiki.org/misc/gray-code/
func circularPermutation(n int, start int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = (i >> 1) ^ i ^ start
	}
	return ans
}
