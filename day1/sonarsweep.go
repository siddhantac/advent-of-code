package main

import "fmt"

var testData = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	243,
}

func main() {
	count := sonarSweepSlidingWindow(depths2)
	fmt.Println(count)
}

func sonarSweepSlidingWindow(depths []int) int {
	sums := []int{}
	for i := 0; ; i++ {
		if i+2 >= len(depths) {
			break
		}
		sum := depths[i] + depths[i+1] + depths[i+2]
		sums = append(sums, sum)
	}
	return sonarsweep1(sums)
}

func sonarsweep1(depths []int) int {
	count := 0

	for i := 0; i < len(depths)-1; i++ {
		if depths[i+1] > depths[i] {
			count++
		}
	}

	return count
}
