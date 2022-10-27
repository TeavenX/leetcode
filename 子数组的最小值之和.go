package main

func main() {

}

const mod int = 1e9 + 7

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	left, right := make([]int, n), make([]int, n)
	stack := make([]int, 0)
	for i, num := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum = (sum + (i-left[i])*(right[i]-i)*arr[i]) % mod
	}
	return sum
}
