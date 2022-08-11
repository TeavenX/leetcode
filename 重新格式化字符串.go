package main

func main() {
	reformatV2("covid2019")
}

func reformat(s string) string {
	count := 0
	nums := make([]rune, 0)
	strs := make([]rune, 0)
	for _, str := range s {
		if str < 'a' {
			count++
			nums = append(nums, str)
		} else {
			count--
			strs = append(strs, str)
		}
	}
	if count > 1 || count < -1 {
		return ""
	}
	if len(nums) > len(strs) {
		nums, strs = strs, nums
	}
	result := make([]rune, len(s))
	for i, str := range strs {
		result[i*2] = str
	}
	for i, num := range nums {
		result[i*2+1] = num
	}
	return string(result)
}

func reformatV2(s string) string {
	count := 0
	result := make([]rune, len(s))
	for _, str := range s {
		if str < 'a' {
			count++
		} else {
			count--
		}
	}
	if count > 1 || count < -1 {
		return ""
	}
	idxnum, idxstr := 0, 1
	if count < 0 {
		idxnum, idxstr = 1, 0
	}
	for _, str := range s {
		if str < 'a' {
			result[idxnum] = str
			idxnum += 2
		} else {
			result[idxstr] = str
			idxstr += 2
		}
	}
	return string(result)
}
