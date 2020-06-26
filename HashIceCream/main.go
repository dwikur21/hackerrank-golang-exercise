package main

import (
	"fmt"
)

func main() {
	arr := []int32{2, 2, 4, 3}
	money := int32(4)

	whatFlavors(arr, money)
}

func whatFlavors(cost []int32, money int32) {
	// Create the hashmap
	iceMap := map[int32][]int32{}

	// Fill the hashmap
	// Key = Icecream Cost, Value = Index in the Cost Slice
	for idx, val := range cost {
		iceMap[val] = append(iceMap[val], int32(idx+1))
	}

	// Looping for each icecream cost
	for _, costEl := range cost {
		val := iceMap[costEl]

		// Search for the remaining money
		x := money - costEl

		// If the remaining money equal to current checked icecream cost
		// e.g. total money = 8, current ice cream cost = 4, then remaining money = 4
		// And there are more than one equal value
		// e.g. there are value of 4 on index 2 and 4
		if x == costEl && len(val) > 1 {

			// Print the result
			fmt.Printf("%v %v", val[0], val[1])

			// Already found the result, then exit loop
			break
		}

		// If the remaining money does not equal to current checked ice cream cost
		if value, found := iceMap[x]; found && x != costEl {

			// Print the result
			fmt.Printf("%v %v", val[0], value[0])

			// Already found the result, then exit loop
			break
		}
	}

	fmt.Println()
}
