package main

func main() {

}

func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, op := range operations {
		switch op[1] {
		case '+':
			result++
		case '-':
			result--
		}
	}
	return result
}
