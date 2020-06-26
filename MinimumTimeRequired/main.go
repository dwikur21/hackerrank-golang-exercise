package main

import (
	"fmt"
	"math"
)

func main() {
	machines := []int64{4, 5, 6}
	goal := int64(12)

	fmt.Println(minTime(machines, goal))
}

func minTime(machines []int64, goal int64) int64 {
	min := min(machines)
	max := max(machines)

	minDay := int64(math.Ceil(float64(goal)/float64(len(machines)))) * min
	maxDay := int64(math.Ceil(float64(goal)/float64(len(machines)))) * max

	for minDay < maxDay {
		day := (minDay + maxDay) / 2
		sum := getSum(machines, day)

		if sum >= goal {
			maxDay = day
		} else {
			minDay = day + 1
		}
	}

	return minDay
}

func min(machines []int64) int64 {
	minimum := int64(0)

	for i, val := range machines {
		if i == 0 || val <= minimum {
			minimum = val
		}
	}
	return minimum
}

func max(machines []int64) int64 {
	maximum := int64(0)

	for i, val := range machines {
		if i == 0 || val >= maximum {
			maximum = val
		}
	}
	return maximum
}

func getSum(machines []int64, day int64) int64 {
	sum := int64(0)

	for _, val := range machines {
		sum += int64(math.Floor(float64(day) / float64(val)))
	}

	return sum
}
