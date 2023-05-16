package main

func main() {

}

func haveConflict(event1 []string, event2 []string) bool {
	st1, et1 := event1[0], event1[1]
	st2, et2 := event2[0], event2[1]
	return !(et1 < st2 || st1 > et2)
}
