package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int32{1, 3, 5, 7, 9}
	k := int32(3)

	fmt.Println(getMinimumCost(k, arr))
}

func getMinimumCost(k int32, c []int32) int32 {
	var count int32
	var result int32

	sort.Slice(c, func(i, j int) bool {
		return c[i] <= c[j]
	})

	for i := len(c) - 1; i >= 0; i-- {
		result += (1 + count/k) * c[i]
		count++
	}

	return result
}
