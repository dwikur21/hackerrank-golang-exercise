package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	a := "fcrxzwscanmligyxyvym"
	b := "jxwtrhvujlmrpdoqbisbwhmgpmeoke"
	// Expected output 30
	fmt.Println(makeAnagram(a, b))
}

func makeAnagram(a string, b string) int32 {
	letterMap := map[string]int32{}
	result := int32(0)

	for _, el := range strings.Split(a, "") {
		letterMap[el]--
	}

	for _, el := range strings.Split(b, "") {
		letterMap[el]++
	}

	for _, el := range letterMap {
		result += int32(math.Abs(float64(el)))
	}

	return result
}
