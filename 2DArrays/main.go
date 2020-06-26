package main

import (
	"fmt"
	"math"
)

func main() {
	numArr := [][]int32{}

	row1 := []int32{3, 7, -3, 0, 1, 8}
	row2 := []int32{1, -4, -7, -8, -6, 5}
	row3 := []int32{-8, 1, 3, 3, 5, 7}
	row4 := []int32{-2, 4, 3, 1, 2, 7}
	row5 := []int32{2, 4, -5, 1, 8, 4}
	row6 := []int32{5, -7, 6, 5, 2, 8}

	numArr = append(numArr, row1)
	numArr = append(numArr, row2)
	numArr = append(numArr, row3)
	numArr = append(numArr, row4)
	numArr = append(numArr, row5)
	numArr = append(numArr, row6)

	result := hourglassSum(numArr)

	fmt.Println(result)
}

func hourglassSum(arr [][]int32) int32 {
	column := len(arr[0])
	row := len(arr)

	sumElement := 0
	sumShape := 0

	for rowShape := 0; rowShape <= row-3; rowShape++ {
		for colShape := 0; colShape <= column-3; colShape++ {
			for insideRow := 0; insideRow < 3; insideRow++ {
				for insideCol := 0; insideCol < 3; insideCol++ {
					if int(math.Floor(float64(insideRow)/float64(2))) <= 1 && (insideRow-insideCol <= 0 && insideRow+insideCol <= 3-1) ||
						int(math.Floor(float64(insideRow)/float64(2))) >= 1 && (insideRow-insideCol >= 0 && insideRow+insideCol >= 3-1) {
						sumElement += int(arr[rowShape+insideRow][colShape+insideCol])
					}
				}
			}

			if (sumElement > sumShape) || (rowShape == 0 && colShape == 0) {
				sumShape = sumElement
			}

			sumElement = 0
		}
	}

	return int32(sumShape)
}
