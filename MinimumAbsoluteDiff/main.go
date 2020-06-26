package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}

	fmt.Println(minimumAbsoluteDifference(arr))
}

func minimumAbsoluteDifference(arr []int32) int32 {
	var result int32
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})

	for i := 0; i < len(arr)-1; i++ {
		if i == 0 || int32(math.Abs(float64(arr[i])-float64(arr[i+1]))) <= result {
			result = int32(math.Abs(float64(arr[i]) - float64(arr[i+1])))
		}
	}

	return result
}
