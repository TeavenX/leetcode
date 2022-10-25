package main

func main() {

}

func shortestSubarray(nums []int, k int) int {
	min := -1
	preSumArr := make([]int, len(nums)+1)
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num
	}
	deq := make([]int, 0)
	for i, presum := range preSumArr {
		// presumArr = [i, j, k, l]
		// 因为有负数存在，因此当presumArr[i] > presumArr[k]时，
		// nums[j]肯定是负数，所以直接看k就好了（因为i的部分已经在前面的窗口就已经计算过 l 了）
		for len(deq) > 0 && preSumArr[deq[len(deq)-1]] > presum {
			deq = deq[:len(deq)-1]
		}
		for len(deq) > 0 && presum-preSumArr[deq[0]] >= k {
			l := i - deq[0]
			if min == -1 || l < min {
				min = l
			}
			deq = deq[1:]
		}
		deq = append(deq, i)
	}
	return min
}
