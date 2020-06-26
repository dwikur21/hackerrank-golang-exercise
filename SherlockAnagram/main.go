package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := "xyyx"
	fmt.Println(sherlockAndAnagrams(str))
}

func sherlockAndAnagrams(s string) int32 {
	strSlice := strings.Split(s, "")
	strMap := map[string]int{}
	var result int32

	for i := 0; i < len(strSlice); i++ {
		substrSlice := []string{}

		for j := i; j < len(strSlice); j++ {
			substrSlice = append(substrSlice, strSlice[j])
			sort.Strings(substrSlice)

			strMap[strings.Join(substrSlice, "")]++
		}
	}

	for _, el := range strMap {
		result += int32((el * (el - 1)) / 2)
	}

	return result
}
