package main

func main() {

}

func longestWPIError(hours []int) int {
	ans, left := 0, 0
	p, n := 0, 0
	for right, hour := range hours {
		if hour <= 8 {
			n++
		} else {
			p++
		}
		for p <= n && left < len(hours) {
			if hours[left] < 8 {
				n--
			} else {
				p--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func longestWPI(hours []int) int {
	s := make([]int, len(hours)+1)
	stack := []int{0}
	for i, hour := range hours {
		i++
		s[i] = s[i-1]
		if hour > 8 {
			s[i]++
		} else {
			s[i]--
		}
		if s[i] < s[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	ans := 0
	for i := len(hours); i > 0; i-- {
		for len(stack) > 0 && s[i] > s[stack[len(stack)-1]] {
			ans = max(ans, i-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}

func longestWPIV2(hours []int) int {
	cache := make([]int, len(hours)+2)
	s := 0
	ans := 0
	for i, h := range hours {
		i++
		if h > 8 {
			s--
		} else {
			s++
		}
		if s < 0 {
			ans = i
		} else {
			// 目标值是s+1
			if cache[s+1] > 0 {
				ans = max(ans, i-cache[s+1])
			}
			// 当前值是s
			if cache[s] == 0 {
				cache[s] = i
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
