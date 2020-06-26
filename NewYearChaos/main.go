package main

import (
	"fmt"
	"math"
)

func main() {
	a1 := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	//a2 := []int32{2, 5, 1, 3, 4}

	minimumBribes(a1)
	//minimumBribes(a2)
}

func minimumBribes(q []int32) {
	bribeCount := int32(0)
	chaotic := false

	for index := int32(0); index < int32(len(q)); index++ {
		if q[index]-(index+1) > 2 {
			chaotic = true
			break
		}

		for j := int32(math.Max(float64(0), float64(q[index]-2))); j < index; j++ {
			if q[j] > q[index] {
				bribeCount++
			}
		}
	}

	if chaotic {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(bribeCount)
	}
}
