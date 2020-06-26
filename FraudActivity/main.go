package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// arr := []int32{10, 20, 30, 40, 50}
	// fmt.Println(activityNotifications(arr, 3))
	fmt.Println(activityNotifications(input200000, 10122)) // Expected Result 1026

	elapsed := time.Since(start)
	fmt.Printf("It took %s to complete", elapsed)
}

func activityNotifications(expenditure []int32, d int32) int32 {
	numNotif := int32(0)
	medianPosition := int32(0)
	current := int32(0)

	if d%2 == 0 {
		medianPosition = d / 2
	} else {
		medianPosition = (d / 2) + 1
	}

	// Create an array(slice) of the size of the maximum value + 1
	k := int32(200)
	count := make([]int32, k+1)

	// Count each element
	for i := int32(0); i < d; i++ {
		count[expenditure[i]] = count[expenditure[i]] + 1
	}

	for i := d; i < int32(len(expenditure)); i++ {
		counter, left := int32(0), int32(0)
		median := float64(0)

		for counter < medianPosition {
			counter += count[left]
			left++
		}

		// Step back one time
		right := left
		left--

		// If odd or both left and right of even are same number
		if counter > medianPosition || d%2 != 0 {
			median = float64(left)
		} else {
			// Find right value for even
			for count[right] == 0 {
				right++
			}
			median = (float64(left) + float64(right)) / 2.0
		}

		if float64(expenditure[i]) >= 2.0*median {
			numNotif++
		}

		count[expenditure[current]]--
		count[expenditure[i]]++

		current++
	}

	return numNotif
}
