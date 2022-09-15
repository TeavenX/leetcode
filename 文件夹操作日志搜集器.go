package main

func main() {

}

func minOperations(logs []string) int {
	depth := 0
	for _, log := range logs {
		switch log {
		case "../":
			depth--
			if depth < 0 {
				depth = 0
			}
		case "./":
			continue
		default:
			depth++
		}
	}
	return depth
}
