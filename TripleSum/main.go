package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int32{1, 3, 5, 7}
	b := []int32{5, 7, 9}
	c := []int32{7, 9, 11, 13}

	fmt.Println(triplets(a, b, c))
}

func triplets(a []int32, b []int32, c []int32) int64 {
	sort.Slice(a, func(i, j int) bool {
		return a[i] <= a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] <= b[j]
	})

	sort.Slice(c, func(i, j int) bool {
		return c[i] <= c[j]
	})

	a = unique(a)
	b = unique(b)
	c = unique(c)

	bprevious := int32(0)
	result := int64(0)

	for _, bVal := range b {

		// aprevious := int32(0)
		// cprevious := int32(0)
		achan := make(chan int64, 0)
		cchan := make(chan int64, 0)

		if bprevious != bVal {
			bprevious = bVal
		} else {
			continue
		}

		go func(achan chan int64) {
			j := int64(0)
			x := sort.Search(len(a), func(i int) bool {
				return a[i] > bVal
			})
			j = int64(x)
			achan <- j
		}(achan)

		go func(cchan chan int64) {
			k := int64(0)
			y := sort.Search(len(c), func(i int) bool {
				return c[i] > bVal
			})
			k = int64(y)
			cchan <- k
		}(cchan)

		left := <-achan
		right := <-cchan
		result += left * right
	}

	return result
}

func unique(intSlice []int32) []int32 {
	keys := make(map[int32]bool)
	list := []int32{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
