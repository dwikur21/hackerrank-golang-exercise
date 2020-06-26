package main

import (
	"fmt"
)

func main() {
	intSlice := []int64{1, 2, 1, 2, 1, 2}
	factor := int64(2)

	fmt.Println(countTriplets(intSlice, factor))
}

func countTriplets(arr []int64, r int64) int64 {
	result := int64(0)

	var mp2 = make(map[int64]int64)
	var mp3 = make(map[int64]int64)

	for _, el := range arr {
		if _, found := mp3[el]; found {
			result += mp3[el]
		}

		if _, found := mp2[el]; found {
			mp3[el*r] += mp2[el]
		}

		mp2[el*r]++
	}

	return result
}
