package main

import (
	"fmt"
	"math"
	"regexp"
)

func main() {
	s := "abaa"
	n := 100008
	if len(s) == 1 && s == "a" {
		fmt.Println(n)
	} else {
		strLengthMultiplier := math.Floor(float64(n) / float64(len(s)))

		r, _ := regexp.Compile("a")

		aCount := len(r.FindAllString(s, -1))

		aCount = aCount * int(strLengthMultiplier)

		if n%len(s) != 0 {
			aCount += len(r.FindAllString(s[:n%len(s)], -1))
		}

		fmt.Println(aCount)
	}
}
