package main

import "fmt"

func main() {
	// magz := []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}
	// note := []string{"ive", "got", "some", "coconuts"}

	magz := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}

	checkMagazine(magz, note)
}

func checkMagazine(magazine []string, note []string) {
	set := make(map[string]int)
	subset := true

	for _, el := range magazine {
		set[el]++
	}

	for _, el := range note {
		if count, found := set[el]; !found {
			subset = false
			break
		} else if count <= 0 {
			subset = false
			break
		} else {
			set[el]--
		}
	}

	if subset {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
