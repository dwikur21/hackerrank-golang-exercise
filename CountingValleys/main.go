package main

import (
	"fmt"
	"strings"
)

func main() {
	countingValleys(8, "UDDDUDUU")
}

func countingValleys(n int32, s string) {
	ss := strings.Split(s, "")

	mountains := int32(0)
	valleys := int32(0)
	altitude := int32(0)

	for i, letter := range ss {
		if letter == "D" {
			altitude--
			if altitude == 0 && i > 0 {
				mountains++
			}
		} else if letter == "U" {
			altitude++
			if altitude == 0 && i > 0 {
				valleys++
			}
		}
	}

	fmt.Println(mountains, valleys)
}
