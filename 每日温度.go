package main

func main() {

}

type Pair struct {
	i, num int
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]Pair, 0)
	for i, num := range temperatures {
		temperatures[i] = 0
		for len(stack) > 0 && stack[len(stack)-1].num < num {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temperatures[p.i] = i - p.i
		}
		stack = append(stack, Pair{i, num})
	}
	return temperatures
}
