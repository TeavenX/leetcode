package main

func main() {

}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum&1 == 1 {
		return []int64{}
	}
	ans := make([]int64, 0)
	var num int64 = 2
	for finalSum > 0 {
		if finalSum < num {
			for i := len(ans) - 1; i >= 0; i-- {
				if ans[i]+finalSum == num {
					ans = append(ans[:i], ans[i+1:]...)
					ans = append(ans, num)
					return ans
				}
			}
		}
		finalSum -= num
		ans = append(ans, num)
		num += 2
	}
	return ans
}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum&1 == 1 {
		return []int64{}
	}
	ans := make([]int64, 0)
	for i := int64(2); i <= finalSum; i += 2 {
		ans = append(ans, i)
		finalSum -= i
	}
	if finalSum > 0 {
		ans[len(ans)-1] += finalSum
	}
	return ans
}
