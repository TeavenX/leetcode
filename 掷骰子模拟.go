package main

func main() {

}

const mod int = 1e9 + 7
const m = 6

func dieSimulatorV1TLE(n int, rollMax []int) int {
	var traceback func(left int, last int, limit int) int
	traceback = func(left int, last int, limit int) int {
		if left == 0 {
			return 1
		}
		ans := 0
		for i := 0; i < m; i++ {
			if i != last {
				ans = (ans + traceback(left-1, i, rollMax[i]-1)) % mod
			} else if limit > 0 {
				ans = (ans + traceback(left-1, last, limit-1)) % mod
			}
		}
		return ans
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = (ans + traceback(n-1, i, rollMax[i]-1)) % mod
	}
	return ans
}

func dieSimulatorV2(n int, rollMax []int) int {
	cache := make(map[int]int)
	var traceback func(left int, last int, limit int) int
	traceback = func(left int, last int, limit int) int {
		if left == 0 {
			return 1
		}
		key := left*1000 + last*20 + limit
		if v, ok := cache[key]; ok {
			return v
		}
		ans := 0
		for i := 0; i < m; i++ {
			if i != last {
				ans = (ans + traceback(left-1, i, rollMax[i]-1)) % mod
			} else if limit > 0 {
				ans = (ans + traceback(left-1, last, limit-1)) % mod
			}
		}
		cache[key] = ans
		return ans
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans = (ans + traceback(n-1, i, rollMax[i]-1)) % mod
	}
	return ans
}

func dieSimulatorV3(n int, rollMax []int) int {
	dp := make([][m][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [m][]int{}
		for j, mx := range rollMax {
			dp[i][j] = make([]int, mx)
		}
	}
	for j := 0; j < m; j++ {
		for k := range dp[0][j] {
			dp[0][j][k] = 1
		}
	}
	for i := 1; i < n; i++ {
		for last, mx0 := range rollMax {
			for left := 0; left < mx0; left++ {
				res := 0
				for j, mx := range rollMax {
					if j != last {
						res += dp[i-1][j][mx-1]
					} else if left > 0 {
						res += dp[i-1][j][left-1]
					}
				}
				dp[i][last][left] = res % mod
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		ans = (ans + dp[n-1][i][rollMax[i]-1]) % mod
	}
	return ans
}

func dieSimulatorCV(n int, rollMax []int) int {
	f := make([][m]int, n)
	for j := range f[0] {
		f[0][j] = 1
	}
	s := make([]int, n)
	s[0] = m
	for i := 1; i < n; i++ {
		for j, mx := range rollMax {
			res := s[i-1]
			if i > mx {
				res -= s[i-mx-1] - f[i-mx-1][j]
			} else if i == mx {
				res--
			}
			f[i][j] = (res%mod + mod) % mod // 防止出现负数
			s[i] = (s[i] + f[i][j]) % mod
		}
	}
	return s[n-1]
}
