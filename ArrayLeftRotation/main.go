package main

import (
	"fmt"
)

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	n := 4

	fmt.Println(rotLeft(arr, int32(n)))

}

func rotLeft(a []int32, d int32) []int32 {
	for count := int32(0); count < d; count++ {
		elTemp := a[0]
		for index := range a {
			if index != len(a)-1 {
				a[index] = a[index+1]
			} else {
				a[index] = elTemp
			}
		}
	}

	return a
}
