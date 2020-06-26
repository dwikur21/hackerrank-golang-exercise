package main

import (
	"fmt"
	"math"
)

func main() {
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	sockColor := map[int32]int32{}

	for _, color := range ar {

		_, exist := sockColor[color]

		if exist {
			sockColor[color] = sockColor[color] + 1
		} else {
			sockColor[color] = 1
		}
	}

	result := int32(0)
	for _, colorCount := range sockColor {
		if colorCount%2 == 0 || colorCount > 1 {
			result += int32(math.Floor(float64(colorCount / 2)))
		}
	}

	fmt.Println(result)
}
