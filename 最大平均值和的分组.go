package main

func main() {

}

/**
nums = [1,2,3,4,5]

[][]arr代表分组的方式，一级索引对应nums的索引，二级缩影对应分割的数组（也就是k）

arr[1][1] = [1]


arr[2][1] = [1,2]
arr[2][2] = [1] + [2]
          = arr[1][1] + [2]


arr[3][1] = [1,2,3]
arr[3][2] = max([1,2] + [3], [1] + [2][3])
          = max(arr[2][1] + [3], arr[1][1] + (preSum[3] - preSum[2]) / (3 - 2))
arr[3][3] = [1] + [2] + [3]
          = arr[2][2] + [3]

arr[i][j] = max(arr[i][j], arr[i'][j-1] + (preSum[i] - preSum[i']) / (i - i')) 其中 i' 属于 [0, i)
*/
func largestSumOfAverages(nums []int, k int) float64 {
	for i, num := range nums {
		if i > 0 {
			num += nums[i-1]
			nums[i] = num
		}
	}
	n := len(nums)
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, k)
		dp[i][0] = float64(nums[i]) / float64(i+1)
	}
	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			for q := 0; q < i; q++ {
				dp[i][j] = max(dp[i][j], dp[q][j-1]+float64(nums[i]-nums[q])/float64(i-q))
			}
		}
	}
	return dp[n-1][k-1]
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
