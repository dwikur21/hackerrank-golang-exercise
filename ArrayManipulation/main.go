package main

import "fmt"

func main() {
	n := int32(5)
	arr := [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}

	fmt.Println(arrayManipulation(n, arr))
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	numList := []int64{}
	maxNum := int64(0)
	tempMax := int64(0)

	for i := int32(0); i < n+1; i++ {
		numList = append(numList, 0)
	}

	for i := int32(0); i < int32(len(queries)); i++ {
		a, b, k := int32(queries[i][0]), int32(queries[i][1]), int32(queries[i][2])

		numList[a] += int64(k)
		if b+int32(1) <= n {
			numList[b+1] -= int64(k)
		}
	}

	for i := int32(1); i <= n; i++ {
		tempMax += numList[i]
		if tempMax > maxNum {
			maxNum = tempMax
		}
	}

	return maxNum
}
