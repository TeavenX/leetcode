package main

func main() {
	prices := []int{8, 4, 6, 2, 3}
	finalPricesV2(prices)
}

func finalPrices(prices []int) []int {
	for i, price := range prices {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= price {
				prices[i] = price - prices[j]
				break
			}
		}
	}
	return prices
}

func finalPricesV2(prices []int) []int {
	// 初始化一个0值的原因是减少后面减去折扣时的判断步骤
	stack := []int{0}
	for i := len(prices) - 1; i >= 0; i-- {
		price := prices[i]
		// 当price小于栈顶元素时，出栈，因为需要跟比自己小的值相减，如果最后是空栈
		// 说明该元素右边不会有比自己更小的值
		for len(stack) > 1 && price < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		prices[i] = price - stack[len(stack)-1]
		// 将当前元素入栈，此时栈中依旧存在一个小于等于price的值
		// 由于上面的for循环会把比price大的值出栈，所以这里入栈也没有影响
		stack = append(stack, price)
	}
	return prices
}
