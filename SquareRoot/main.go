package main

import "fmt"

// Sqrt : Square Root function
func Sqrt(x float64) {
	for index := 1; index < 11; index++ {
		z := float64(index)
		for j := 0; j < 10; j++ {
			z -= (z*z - x) / (2 * z)
			fmt.Println(z)
		}
	}
}

func main() {
	Sqrt(float64(4))
	// fmt.Println(Sqrt(2))
}
