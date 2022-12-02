package main

func main() {

}

// func minOperations(boxes string) []int {
//     result := make([]int, len(boxes))
//     cache := make(map[int]struct{})
//     for i, box := range boxes {
//         if box == '1' {
//             cache[i] = struct{}{}
//         }
//     }
//     for i := range boxes {
//         t := 0
//         for j, _ := range cache {
//             t += abs(j-i)
//         }
//         result[i] = t
//     }
//     return result
// }

// func abs(a int) int {
//     if a < 0 {
//         return -a
//     }
//     return a
// }

// func minOperations(boxes string) []int {
//     n := len(boxes)
//     result := make([]int, n)
//     for i, box := range boxes {
//         if box == '1' {
//             for j := 0; j < n; j++ {
//                 result[j] += abs(j-i)
//             }
//         }
//     }
//     return result
// }

func minOperations(boxes string) []int {
	lc, rc := 0, 0
	n := len(boxes)
	result := make([]int, n)
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			rc++
			result[0] += i
		}
	}
	if boxes[0] == '1' {
		lc++
	}
	for i := 1; i < n; i++ {
		result[i] = result[i-1] + lc - rc
		if boxes[i] == '1' {
			lc++
			rc--
		}
	}
	return result
}
