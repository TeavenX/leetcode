package main

func main() {

}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, b := range s {
		stack = append(stack, b)
		if len(stack) >= 3 && string(stack[len(stack)-3:]) == "abc" {
			stack = stack[:len(stack)-3]
		}
	}
	return len(stack) == 0
}
