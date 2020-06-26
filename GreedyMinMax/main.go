package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int32{4504, 1520, 5857, 4094, 4157, 3902, 822, 6643, 2422, 7288, 8245, 9948, 2822, 1784, 7802, 3142, 9739, 5629, 5413, 7232}
	k := int32(5)

	fmt.Println(maxMin(k, arr)) // 1335
}

func maxMin(k int32, arr []int32) int32 {
	var minResult int32
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})

	minResult = arr[k-1] - arr[0]
	for i := 0; i <= len(arr)-int(k); i++ {
		temp := arr[i+(int(k)-1)] - arr[i]

		if temp <= minResult {
			minResult = temp
		}
	}

	return minResult
}
