package main

import "fmt"

func main() {
	//asteroids := []int{5, 10, -5}
	//asteroids := []int{5, 10, -10}
	asteroids := []int{-2, -1, 1, 2}
	fmt.Println(asteroidCollision(asteroids))
}

func asteroidCollision(asteroids []int) []int {
	result := make([]int, 0)
next:
	for _, asteroid := range asteroids {
		if len(result) == 0 {
			result = append(result, asteroid)
		} else {
			for len(result) > 0 && result[len(result)-1] > 0 && result[len(result)-1]*asteroid < 0 {
				if abs(result[len(result)-1]) < abs(asteroid) {
					result = result[:len(result)-1]
				} else if abs(result[len(result)-1]) == abs(asteroid) {
					result = result[:len(result)-1]
					continue next
				} else {
					continue next
				}
			}
			result = append(result, asteroid)
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
