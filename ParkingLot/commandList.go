package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var parkingArea []ParkingLot
var emptyParkingLot []int

func createParkingLot(args []string) {
	if len(args) > 0 {
		num, err := strconv.Atoi(args[0])
		check(err)

		for i := 0; i < num; i++ {
			parkingArea = append(parkingArea, ParkingLot{lotNum: i + 1, car: Car{}})
			emptyParkingLot = append(emptyParkingLot, i+1)
		}

		fmt.Printf("Created a parking lot with %v slots \n", len(parkingArea))
	} else {
		fmt.Printf("Please insert how many parking lot(s)\n")
	}
}

func park(args []string) {
	if len(args) > 0 {
		licensePlate := args[0]
		carColour := args[1]
		lotNum := -1

		if len(emptyParkingLot) > 0 {
			sort.Ints(emptyParkingLot)
			emptyParkingIdx := emptyParkingLot[0]
			if len(emptyParkingLot) > 1 {
				emptyParkingLot = emptyParkingLot[1:]
			} else {
				emptyParkingLot = nil
			}

			parkingArea[emptyParkingIdx-1].car = Car{colour: carColour, plateNum: licensePlate}
			lotNum = parkingArea[emptyParkingIdx-1].lotNum

			fmt.Printf("Allocated slot number: %v \n", lotNum)
		} else {
			fmt.Printf("Sorry, parking lot is full \n")
		}
	} else {
		fmt.Printf("Please insert the car's license number and colour\n")
	}

}

func leave(args []string) {
	if len(args) > 0 {
		lotNum, err := strconv.Atoi(args[0])
		check(err)

		parkingArea[lotNum-1].car = Car{}
		emptyParkingLot = append(emptyParkingLot, lotNum)

		fmt.Printf("Slot number %v is free \n", lotNum)
	} else {
		fmt.Printf("Please insert the car's license number\n")
	}
}

func status() {
	if len(parkingArea) > 0 {
		fmt.Printf("Slot No.\tRegistration No\t\tColour\n")

		for _, val := range parkingArea {
			if val.car != (Car{}) {
				fmt.Printf("%v\t\t%v\t\t%v\n", val.lotNum, val.car.plateNum, val.car.colour)
			}
		}
	} else {
		fmt.Printf("Parking lot is empty\n")
	}
}

func registrationNumbersForCarsWithColour(args []string) {
	if len(args) > 0 {
		colour := args[0]
		arrRes := []string{}

		for _, val := range parkingArea {
			if val.car.colour == colour {
				arrRes = append(arrRes, val.car.plateNum)
			}
		}

		resString := strings.Join(arrRes, ", ")

		fmt.Printf("%v\n", resString)
	} else {
		fmt.Printf("Please insert the car's colour\n")
	}
}

func slotNumbersForCarsWithColour(args []string) {
	if len(args) > 0 {
		colour := args[0]
		arrRes := []string{}

		for _, val := range parkingArea {
			if val.car.colour == colour {
				arrRes = append(arrRes, strconv.Itoa(val.lotNum))
			}
		}

		resString := strings.Join(arrRes, ", ")

		fmt.Printf("%v\n", resString)
	} else {
		fmt.Printf("Please insert the car's colour\n")
	}
}

func slotNumberForRegistrationNumber(args []string) {
	if len(args) > 0 {
		licensePlate := args[0]
		found := false
		lotNum := -1

		for _, val := range parkingArea {
			if val.car.plateNum == licensePlate {
				found = true
				lotNum = val.lotNum
				break
			}
		}

		if found {
			fmt.Printf("%v\n", lotNum)
		} else {
			fmt.Printf("Not found\n")
		}
	} else {
		fmt.Printf("Please insert the car's license number\n")
	}
}
