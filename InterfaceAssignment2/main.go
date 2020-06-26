package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]

	file, err := os.Open(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 1024)
	file.Read(buff)

	fmt.Println(string(buff))
}
