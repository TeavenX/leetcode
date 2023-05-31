package main

func main() {

}

func minRemoveToMakeValid(s string) string {
	bf := make([]rune, 0, len(s))
	stack := make([]int, 0)
	for _, b := range s {
		if b == ')' {
			if len(stack) == 0 {
				continue
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if b == '(' {
			stack = append(stack, len(bf))
		}
		bf = append(bf, b)
	}
	for i := len(stack) - 1; i >= 0; i-- {
		idx := stack[i]
		bf = append(bf[:idx], bf[idx+1:]...)
	}
	return string(bf)
}
