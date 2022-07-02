package main

func main() {

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	left, right := 0, n
	for left < right {
		mid := (right-left)>>1 + left
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
