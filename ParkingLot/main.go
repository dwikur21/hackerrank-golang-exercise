package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1]

		cmdArr := readFile(args)

		for _, val := range cmdArr {
			command(val)
		}

		mainApp()
	} else {
		mainApp()
	}
}

func check(e error) {
	if e != nil {
		fmt.Printf(e.Error())
	}
}

func mainApp() {
	reader := bufio.NewReader(os.Stdin)
	for {
		commandArgs, err := reader.ReadString('\n')
		check(err)

		commandArgs = strings.Replace(commandArgs, "\r\n", "", -1)
		command(commandArgs)
	}
}

func readFile(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	check(err)

	strCmd := string(data)
	result := strings.Split(strCmd, "\r\n")

	return result
}

func command(cmd string) {
	cmdArr := strings.Split(cmd, " ")
	args := cmdArr[1:]

	switch cmdArr[0] {
	case "create_parking_lot":
		createParkingLot(args)
	case "park":
		park(args)
	case "leave":
		leave(args)
	case "status":
		status()
	case "registration_numbers_for_cars_with_colour":
		registrationNumbersForCarsWithColour(args)
	case "slot_numbers_for_cars_with_colour":
		slotNumbersForCarsWithColour(args)
	case "slot_number_for_registration_number":
		slotNumberForRegistrationNumber(args)
	default:
		fmt.Printf("Command not recognized\n")
	}
}
