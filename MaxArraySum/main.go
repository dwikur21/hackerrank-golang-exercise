package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int32{3, 5, -7, 8, 10}

	res := maxSubsetSum(arr)

	fmt.Println(res)
}

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
	// Check if the array contains any element
	// The sum of 0 subset is 0
	if len(arr) == 0 {
		return int32(0)
	}

	// Check if the array contains only 1 element
	// If yes, return the first element
	arr[0] = int32(math.Max(float64(0), float64(arr[0])))
	if len(arr) == 1 {
		return arr[0]
	}

	// At index 1, the max value could be at index 0 or
	// at index 1 itself
	arr[1] = int32(math.Max(float64(arr[0]), float64(arr[1])))

	// Loop for the length of the array but start at index 2
	for index := 2; index < len(arr); index++ {

		// At index 2, calculate which is greater
		// Previous value at current index - 1 or current value plus previous value at current index - 2
		// That way, at each iteration we keep track of current max value that are stored at current index - 1
		arr[index] = int32(math.Max(float64(arr[index-1]), float64(arr[index-2]+arr[index])))
	}

	// After the sum, we keep the max value at last index
	// Return it
	return arr[len(arr)-1]
}
