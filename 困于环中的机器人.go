package main

func main() {

}

var steps = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	idx := 0
	for _, b := range instructions {
		switch b {
		case 'G':
			step := steps[(idx+4)%4]
			x, y = x+step[0], y+step[1]
		case 'L':
			idx = (idx - 1) % 4
		case 'R':
			idx = (idx + 1) % 4
		}
	}
	// 因为是循环走，一轮走完存在一下几种情况：
	// 1、坐标回到(0,0)，且方向朝北 -- 最终结果一定是困在循环中
	// 2、坐标为(x,y)，且方向朝北 -- 下一次循环，坐标会变为 (2x,2y)，一定不会困在循环中
	// 3、坐标为(x,y)，且方向朝南 -- 下一次循环会-x，-y，回到(0,0)且方向朝北，困在循环中
	// 4、如果方向朝东/西，也就是转了90度，那么再走3次（一共走了4次）之后，会回到原点，且方向朝北 -- 困在循环中
	// 结论：
	// 已经回到(0, 0) 或者 方向朝北 的情况不会困在循环中
	return idx != 0 || x == 0 && y == 0
}
