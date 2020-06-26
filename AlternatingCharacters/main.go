package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "AAAAA"

	fmt.Println(alternatingCharacters(a))
}

func alternatingCharacters(s string) int32 {
	stringArr := strings.Split(s, "")
	temp := stringArr[0]
	count := int32(0)

	for i := 1; i < len(stringArr); i++ {
		if stringArr[i] == temp {
			count++
		}
		temp = stringArr[i]
	}

	return count
}
