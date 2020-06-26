package main

import (
	"fmt"
	"sort"
)

type priceList []int32

func (a priceList) Len() int           { return len(a) }
func (a priceList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a priceList) Less(i, j int) bool { return a[i] <= a[j] }

func main() {
	arr := []int32{1, 12, 5, 111, 200, 1000, 10}

	fmt.Println(maximumToys(arr, 50))
}

func maximumToys(prices []int32, k int32) int32 {
	toysPriceList := priceList(prices)

	sort.Sort(toysPriceList)

	sumVal := int32(0)
	numOfToys := int32(0)
	for i, val := range toysPriceList {
		if sumVal+val > k {
			numOfToys = int32(i)
			break
		}
		sumVal += val
	}

	return numOfToys
}
