package main

func main() {

}

// func secondHighest(s string) int {
//     cache := [10]bool{}
//     for _, str := range s {
//         if str >= '0' && str <= '9' {
//             num := int(str - '0')
//             cache[num] = true
//         }
//     }
//     flag := false
//     for i := 9; i >= 0; i-- {
//         if flag && cache[i] {
//             return i
//         }
//         flag = flag || cache[i]
//     }
//     return -1
// }

// func secondHighest(s string) int {
//     var cache int = 0
//     for _, str := range s {
//         if str >= '0' && str <= '9' {
//             num := str - '0'
//             cache |= 1 << num
//         }
//     }
//     flag := 0
//     for i := 9; i >= 0; i-- {
//         num := (cache >> i) & 1
//         if flag == 1 && num == 1 {
//             return i
//         }
//         flag = flag | num
//     }
//     return -1
// }

func secondHighest(s string) int {
	first, second := -1, -1
	for _, str := range s {
		if str >= '0' && str <= '9' {
			num := int(str - '0')
			if num > first {
				second = first
				first = num
			} else if num > second && num < first {
				second = num
			}
		}
	}
	return second
}

// func secondHighest(s string) int {
//     first, second := -1, -1
//     for _, str := range s {
//         if unicode.IsDigit(str) {
//             num := int(str - '0')
//             if num > first {
//                 second = first
//                 first = num
//             } else if num > second && num < first {
//                 second = num
//             }
//         }
//     }
//     return second
// }
