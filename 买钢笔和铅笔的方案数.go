package main

func main() {

}

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ans int64 = 0
	for i := 0; cost1*i <= total; i++ {
		res := (total-cost1*i)/cost2 + 1
		ans += int64(res)
	}
	return ans
}
