package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

}

func coinChangeError(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount)
	for _, coin := range coins {
		for j := amount - 1; j >= coin; j-- {
			temp := dp[j-coin] + coin
			if temp > dp[j] {
				dp[j] = temp
			}
		}
	}
	if dp[amount-1] != amount {
		return -1
	}
	count, n := 0, len(coins)
	sort.Ints(coins)
	for amount > 0 {
		coin := coins[n]
		if coin > amount {
			n--
		} else {
			amount -= coin
			count++
		}
	}
	return count
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt
		for _, coin := range coins {
			if j >= coin && dp[j-coin]+1 < dp[j] {
				dp[j] = dp[j-coin] + 1
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func coinChange20220512(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = 10e4
		for _, coin := range coins {
			if coin <= i {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
