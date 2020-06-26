package main

import (
	"fmt"
	"sort"
)

func main() {
	k := int32(687)
	//arr := []int32{1, 5, 3, 4, 2}

	result := pairs(k, arr)

	fmt.Println(result) // Expected 80
}

func pairs(k int32, arr []int32) int32 {
	// Return variable
	count := int32(0)

	// Creating variable for the running index
	// It is used for two pointer technique
	i, j := int32(0), int32(1)

	// Sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})

	// Loop the logic as long as j is less than
	// the length of the array, so it doesn't throw
	// an error of array index out of bounds
	for j < int32(len(arr)) {

		// If i equals to j, increment j for the next value
		// Because there is no meaning comparing the same value
		if i == j {
			j++
			continue
		}

		// Calculate the difference
		diff := arr[j] - arr[i]

		// If the difference is equal to query
		// then increment the count that fulfilled the condition
		if diff == k {
			count++

			// Increment the j index, for comparing the next value
			j++
		} else if diff > k {

			// If the difference is bigger than the query
			// Increment the i index
			i++
		} else if diff < k {

			// If the difference is less than the query
			// Increment the j index
			j++
		}
	}

	return count
}
