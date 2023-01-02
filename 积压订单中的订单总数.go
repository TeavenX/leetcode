package main

import "container/heap"

func main() {

}

type pair struct {
	price, amount int
}

type hp []pair

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i].price < h[j].price }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }

func (h *hp) Pop() interface{} {
	t := *h
	v := t[len(t)-1]
	t = t[:len(t)-1]
	*h = t
	return v
}

const mod int = 1e9 + 7

func getNumberOfBacklogOrders(orders [][]int) int {
	sell, buy := hp{}, hp{}
	for _, order := range orders {
		price := order[0]
		amount := order[1]
		if order[2] == 0 {
			// buy
			for amount > 0 && len(sell) > 0 && sell[0].price <= price {
				b := heap.Pop(&sell).(pair)
				if amount >= b.amount {
					amount -= b.amount
				} else {
					heap.Push(&sell, pair{price: b.price, amount: b.amount - amount})
					amount = 0
				}
			}
			if amount > 0 {
				heap.Push(&buy, pair{price: -price, amount: amount})
			}
		} else {
			// sell
			for amount > 0 && len(buy) > 0 && -buy[0].price >= price {
				s := heap.Pop(&buy).(pair)
				if amount >= s.amount {
					amount -= s.amount
				} else {
					heap.Push(&buy, pair{price: s.price, amount: s.amount - amount})
					amount = 0
				}
			}
			if amount > 0 {
				heap.Push(&sell, pair{price: price, amount: amount})
			}
		}
	}
	sum := 0
	for len(sell) > 0 {
		sum += heap.Pop(&sell).(pair).amount
	}
	for len(buy) > 0 {
		sum += heap.Pop(&buy).(pair).amount
	}
	return sum % mod
}
