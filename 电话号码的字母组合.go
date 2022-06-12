package main

func main() {

}

var digiMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	n := len(digits)
	if n == 0 {
		return result
	}
	path := ""
	var traceback func(startIdx int)
	traceback = func(startIdx int) {
		if startIdx == n {
			result = append(result, path)
			return
		}
		num := digits[startIdx]
		vabList := digiMap[num]
		for i := 0; i < len(vabList); i++ {
			path += vabList[i]
			traceback(startIdx + 1)
			path = path[:len(path)-1]
		}
	}
	traceback(0)
	return result
}
