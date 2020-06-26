package main

import "fmt"

func main() {
	c := []int32{0, 0, 0, 1, 0, 0}
	//c := []int32{0, 1, 0, 1, 0, 1, 0}
	//c := []int32{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}

	steps := int32(0)
	index := int(0)

	for ok := true; ok; ok = index < len(c)-1 {
		if index < len(c)-2 {
			if c[index+2] != 1 {
				steps++
				index += 2
				continue
			}
		}

		if c[index+1] != 1 {
			steps++
			index++
			continue
		}
	}

	fmt.Println(steps)
}
