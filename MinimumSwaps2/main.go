package main

import "fmt"

func main() {
	arr := []int32{1, 3, 5, 2, 4, 6, 7}

	result := minimumSwaps(arr)

	fmt.Println(result)
}

func minimumSwaps(arr []int32) int32 {
	swapCount := int32(0)

	for i, el := range arr {
		if el == int32(i+1) {
			continue
		} else {
			for j := int32(i + 1); j < int32(len(arr)); j++ {
				if arr[j] == int32(i+1) {
					arr[i] = arr[j]
					arr[j] = el

					swapCount++
				}
			}
		}
	}

	return swapCount
}
