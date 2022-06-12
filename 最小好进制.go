package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

func main() {
	fmt.Println(smallestGoodBase("13"))
}

func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	mMax := bits.Len(uint(nVal)) - 1 // bits.Len 用来获取表示目标数所需要的最小位数，输入3输出2（3 = 11）
	// 通过 bits.Len来限制枚举的深度
	for m := mMax; m > 1; m-- {
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		// 开m次方根
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			// 模拟 k^0 + k^1 + k^2 + ... + k^m-1
			mul *= k
			sum += mul
		}
		if sum == nVal {
			// 枚举完一个k，校验一遍是否刚好等于输入的数值
			// 如果等于直接返回结果
			return strconv.Itoa(k)
		}
	}
	// 如果所有的k枚举完都不匹配，就返回最差的结果，也就是m进制表示结果为`11`的情况，此时 m进制的 10 = n - 1；11 = n - 1 + 1 = n
	return strconv.Itoa(nVal - 1)
}
