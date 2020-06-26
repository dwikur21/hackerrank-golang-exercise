package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(dayOfProgrammer(1800))
}

func dayOfProgrammer(year int32) string {
	arrDay := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	month := 0
	numDay := 0

	if year <= 1917 {
		i := 0
		for numDay+arrDay[i] < 256 && i < len(arrDay) {
			if i == 1 {
				if year%4 == 0 {
					numDay += arrDay[i] + 1
					month++
				} else {
					numDay += arrDay[i]
					month++
				}
			} else {
				numDay += arrDay[i]
				month++
			}

			i++
		}
	} else if year == 1918 {
		i := 0
		for numDay+arrDay[i] < 256 && i < len(arrDay) {
			if i == 1 {
				numDay += 15
				month++
			} else {
				numDay += arrDay[i]
				month++
			}

			i++
		}
	} else {
		i := 0
		for numDay+arrDay[i] < 256 && i < len(arrDay) {
			if i == 1 {
				if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
					numDay += arrDay[i] + 1
					month++
				} else {
					numDay += arrDay[i]
					month++
				}
			} else {
				numDay += arrDay[i]
				month++
			}

			i++
		}
	}

	stringNum := ""
	if (256 - numDay) < 10 {
		stringNum = "0" + strconv.Itoa(256-numDay)
	} else {
		stringNum = strconv.Itoa(256 - numDay)
	}

	monthNum := ""
	if (month + 1) < 10 {
		monthNum = "0" + strconv.Itoa(month+1)
	} else {
		monthNum = strconv.Itoa(month + 1)
	}

	return fmt.Sprintf("%v.%v.%v", stringNum, monthNum, year)
}
