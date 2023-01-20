package main

func main() {

}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	cache := make(map[int]map[int]struct{})
	result := make([]int, k)
	for _, log := range logs {
		id, time := log[0], log[1]
		if _, ok := cache[id]; !ok {
			cache[id] = make(map[int]struct{})
		}
		cache[id][time] = struct{}{}
	}
	for _, v := range cache {
		result[len(v)-1]++
	}
	return result
}
