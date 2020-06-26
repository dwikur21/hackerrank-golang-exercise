package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello"
	s2 := "world"

	fmt.Println(twoStrings(s1, s2))
}

func twoStrings(s1 string, s2 string) string {
	a := strings.Split(s1, "")
	b := strings.Split(s2, "")
	issubstring := false

	set := make(map[string]int)
	for _, el := range a {
		set[el]++
	}

	for _, val := range b {
		if _, found := set[val]; !found {
			continue
		} else {
			issubstring = true
			set[val]--
		}
	}

	if issubstring {
		return "YES"
	}
	return "NO"
}
