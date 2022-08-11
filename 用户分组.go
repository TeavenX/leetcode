package main

func main() {

}

func groupThePeople(groupSizes []int) [][]int {
	result := make([][]int, 0)
	cache := make(map[int][]int)
	for i, size := range groupSizes {
		if size == 1 {
			result = append(result, []int{i})
		} else {
			cache[size] = append(cache[size], i)
		}
	}
	for k, v := range cache {
		temp := make([]int, 0)
		for _, item := range v {
			temp = append(temp, item)
			if len(temp) == k {
				result = append(result, append([]int{}, temp...))
				temp = temp[0:0:0]
			}
		}
	}
	return result
}

func groupThePeopleV2(groupSizes []int) [][]int {
	result := make([][]int, 0)
	cache := make(map[int][]int)
	for i, size := range groupSizes {
		if size == 1 {
			result = append(result, []int{i})
		} else {
			cache[size] = append(cache[size], i)
		}
	}
	for k, v := range cache {
		for i := 0; i < len(v); i += k {
			result = append(result, v[i:i+k])
		}
	}
	return result
}
