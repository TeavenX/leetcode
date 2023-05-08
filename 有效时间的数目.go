package main

func main() {

}

func calc(t string, time int) int {
	ans := 0
	for i := 0; i < time; i++ {
		if (t[0] == '?' || i/10 == int(t[0]-'0')) && (t[1] == '?' || i%10 == int(t[1]-'0')) {
			ans++
		}
	}
	return ans
}

func countTime(time string) int {
	return calc(time[:2], 24) * calc(time[3:], 60)
}
