package main

func main() {

}

func generateParenthesisError(n int) []string {
	var quoteAdd = [][]string{{"()", ""}, {"", "()"}, {"(", ")"}}
	cache := make(map[string]bool)
	result := make([]string, 0)
	path := ""
	var traceback func(quoteNum int)
	traceback = func(quoteNum int) {
		if quoteNum == n {
			result = append(result, path)
			return
		}
		for i := 0; i < len(quoteAdd); i++ {
			temp := path
			path = quoteAdd[i][0] + path
			path += quoteAdd[i][1]
			if exist := cache[path]; !exist {
				traceback(quoteNum + 1)
				cache[path] = true
			}
			path = temp
		}
	}
	traceback(0)
	return result
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	var traceback func(path string, left, right int)
	traceback = func(path string, left, right int) {
		if left == n && right == n {
			result = append(result, path)
			return
		}
		if right > left {
			return
		}
		if left < n {
			traceback(path+"(", left+1, right)
		}
		if right < n {
			traceback(path+")", left, right+1)
		}
	}
	traceback("", 0, 0)
	return result
}
