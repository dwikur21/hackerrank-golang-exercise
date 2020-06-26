package main

import (
	"fmt"
	"sort"
)

func main() {
	contests := [][]int32{
		{13, 1},
		{10, 1},
		{9, 1},
		{8, 1},
		{13, 1},
		{12, 1},
		{18, 1},
		{13, 1},
	}
	k := int32(5)

	fmt.Println(luckBalance(k, contests)) //expect 42
}

func luckBalance(k int32, contests [][]int32) int32 {
	sort.Slice(contests, func(i, j int) bool {
		return contests[i][1] <= contests[j][1]
	})

	sort.Slice(contests, func(i, j int) bool {
		return contests[i][0] >= contests[j][0]
	})

	var tempResult int32
	var offset int32
	for j, element := range contests {
		if element[1] == 0 {
			tempResult += element[0]
			offset++
		} else {
			if int32(j) < k+offset {
				tempResult += element[0]
			} else {
				tempResult -= element[0]
			}
		}
	}

	//fmt.Println(contests, tempResult)

	return tempResult
}
