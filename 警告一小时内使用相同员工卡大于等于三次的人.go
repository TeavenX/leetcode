package main

import "sort"

func main() {

}

func alertNames(keyName []string, keyTime []string) []string {
	keyTimeMap := make(map[string][]int)
	for i, name := range keyName {
		time := keyTime[i]
		hour := int(time[0]-'0')*10 + int(time[1]-'0')
		minute := int(time[3]-'0')*10 + int(time[4]-'0')
		keyTimeMap[name] = append(keyTimeMap[name], hour*60+minute)
	}
	ans := make([]string, 0)
	for name, v := range keyTimeMap {
		sort.Ints(v)
		for i := 2; i < len(v); i++ {
			if v[i]-v[i-2] <= 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	sort.Strings(ans)
	return ans
}
