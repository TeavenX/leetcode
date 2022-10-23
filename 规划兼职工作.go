package main

import (
	"fmt"
	"sort"
)

func main() {
	starTime := []int{24, 24, 7, 2, 1, 13, 6, 14, 18, 24}
	endTime := []int{27, 27, 20, 7, 14, 22, 20, 24, 19, 27}
	profit := []int{6, 1, 4, 2, 3, 6, 5, 6, 9, 8}
	fmt.Println(jobScheduling(starTime, endTime, profit))
}

func jobSchedulingError(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		return endTime[index[i]] < endTime[index[j]]
	})
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		idx := index[i]
		j := sort.Search(idx, func(j int) bool { return endTime[index[j]] > startTime[idx] })
		dp[i+1] = max(dp[i], dp[j]+profit[idx])
	}
	fmt.Println(dp)
	return dp[n]
}

type tuple struct {
	start, end int
	job        int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]tuple, n)
	for i := range jobs {
		jobs[i] = tuple{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})
	dp := make([]int, n+1)
	for i, job := range jobs {
		j := sort.Search(i, func(j int) bool { return jobs[j].end > job.start })
		dp[i+1] = max(dp[i], dp[j]+job.job)
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
