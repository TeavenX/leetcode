package main

func main() {

}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	d1, d2 := 0, 0
	n := len(distance)
	for i := start; i%n != destination; i++ {
		d1 += distance[i%n]
	}
	for i := start; (n+i)%n != destination; i-- {
		d2 += distance[(n+i-1)%n]
	}
	if d1 < d2 {
		return d1
	}
	return d2
}

func distanceBetweenBusStopsV2(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	sum1, sum2 := 0, 0
	for i, d := range distance {
		if i >= start && i < destination {
			sum1 += d
		} else {
			sum2 += d
		}
	}
	if sum1 < sum2 {
		return sum1
	}
	return sum2
}
