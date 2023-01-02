package main

import "fmt"

func main() {
	// orders := [][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}
	// orders := [][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}
	orders := [][]int{{1, 29, 1}, {22, 7, 1}, {24, 1, 0}, {25, 15, 1}, {18, 8, 1}, {8, 22, 0}, {25, 15, 1}, {30, 1, 1}, {27, 30, 0}}
	fmt.Println(getNumberOfBacklogOrders(orders))
}

const mod int = 1e9 + 7

type tuple struct {
	preMax int
	preMin int
	amount int
}

func getNumberOfBacklogOrders(orders [][]int) int {
	buyCache := make(map[int]tuple)
	sellCache := make(map[int]tuple)
	var (
		minSell int = 1e9
		maxSell int = 1
		minBuy  int = 1e9
		maxBuy  int = 1
	)
	for _, order := range orders {
		amount := order[1]
		price := order[0]
		if order[2] == 0 {
			// buy
			for i := minSell; i <= price && amount > 0; i++ {
				if sell, ok := sellCache[i]; ok {
					if sell.amount > amount {
						sell.amount -= amount
						sellCache[i] = sell
						amount = 0
					} else if sell.amount > 0 {
						amount -= sell.amount
						maxSell = sell.preMax
						minSell = sell.preMin
						delete(sellCache, i)
					}
				}
			}
			if amount > 0 {
				if c, ok := buyCache[price]; ok {
					c.amount += amount
					buyCache[price] = c
				} else {
					buyCache[price] = tuple{amount: amount, preMax: maxBuy, preMin: minBuy}
				}
				if price > maxBuy {
					maxBuy = price
				}
				if price < minBuy {
					minBuy = price
				}
			}
		} else {
			// sell
			for i := maxBuy; i >= price && amount > 0; i-- {
				if buy, ok := buyCache[i]; ok {
					if buy.amount > amount {
						buy.amount -= amount
						buyCache[i] = buy
						amount = 0
					} else if buy.amount > 0 {
						amount -= buy.amount
						maxBuy = buy.preMax
						minBuy = buy.preMin
						delete(buyCache, i)
					}
				}
			}
			if amount > 0 {
				if c, ok := sellCache[price]; ok {
					c.amount += amount
					sellCache[price] = c
				} else {
					sellCache[price] = tuple{amount: amount, preMax: maxSell, preMin: minSell}
				}
				if price > maxSell {
					maxSell = price
				}
				if price < minSell {
					minSell = price
				}
			}
		}
	}
	sum := 0
	for _, v := range buyCache {
		sum = (sum + v.amount) % mod
	}
	for _, v := range sellCache {
		sum = (sum + v.amount) % mod
	}
	return sum
}
