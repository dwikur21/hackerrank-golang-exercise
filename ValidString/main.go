package main

import (
	"fmt"
	"strings"
)

func main() {
	//a := "aabbccddeefghi" //Expected "NO"
	//a := "abcdefghhgfedecba" //Expected "YES"
	a := "abcdefghhgfedecba" //Expected "YES"

	fmt.Println(isValid(a))
}

func isValid(s string) string {
	letterMap := map[string]int{}
	count := int64(0)
	countDiff := int64(-1)
	valid := "YES"

	for _, el := range strings.Split(s, "") {
		letterMap[el]++
	}

	for _, el := range letterMap {
		count = int64(el)
		break
	}

	for _, el := range letterMap {
		if count != int64(el) {
			countDiff++
		}

		if countDiff >= 1 {
			valid = "NO"
			break
		}
	}

	return valid
}
