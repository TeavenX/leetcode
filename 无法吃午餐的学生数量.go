package main

func main() {

}

func countStudents(students []int, sandwiches []int) int {
	i := 0
	n := len(students)
	cache := [2]int{}
	for _, stu := range students {
		cache[stu]++
	}
	for {
		idx := i % n
		for students[idx] != -1 && students[idx] == sandwiches[0] {
			cache[students[idx]]--
			students[idx] = -1
			sandwiches = sandwiches[1:]
			i++
		}
		i++
		if len(sandwiches) == 0 || cache[sandwiches[0]] == 0 {
			break
		}
	}
	return len(sandwiches)
}
