package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// arr := [][]int32{
	// 	{1, 3},
	// 	{2, 3},
	// 	{3, 2},
	// 	{1, 4},
	// 	{1, 5},
	// 	{1, 5},
	// 	{1, 4},
	// 	{3, 2},
	// 	{2, 4},
	// 	{3, 2},
	// }

	// for _, el := range freqQuery(arr) {
	// 	fmt.Println(el)
	// }
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0777)
	checkError(err)
	reader := bufio.NewReaderSize(file, 16*1024*1024)

	resultFile, err := os.OpenFile("result.txt", os.O_RDONLY, 0777)
	checkError(err)
	readerResult := bufio.NewReaderSize(resultFile, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	qResult, err := strconv.ParseInt(strings.TrimSpace(readLine(readerResult)), 10, 64)
	checkError(err)
	r := int32(qResult)
	var result []int32
	for i := 0; i < int(r); i++ {
		qResultTemp := strings.TrimRight(readLine(readerResult), " \t\r\n")

		queriesItemTemp, err := strconv.ParseInt(qResultTemp, 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		result = append(result, queriesItem)
	}

	ans := freqQuery(queries)
	if len(ans) == len(result) {
		for i := 0; i < len(result); i++ {
			if ans[i] == result[i] {
				continue
			} else {
				fmt.Println(i, ans[i], result[i])
			}
		}
	} else {
		fmt.Println(false)
	}

}

func freqQuery(queries [][]int32) []int32 {
	result := []int32{}

	mapFreq := map[int64]int64{}
	mapFreqOfFreq := map[int64]int64{}

	for i := 0; i < len(queries); i++ {
		el := queries[i]

		if el[0] == 1 {
			mapFreq[int64(el[1])]++

			mapFreqOfFreq[mapFreq[int64(el[1])]]++
			if val, found := mapFreqOfFreq[mapFreq[int64(el[1])]-1]; found && val > 0 {
				mapFreqOfFreq[mapFreq[int64(el[1])]-1]--
			}

		} else if el[0] == 2 {
			if mapFreq[int64(el[1])] > 0 {
				mapFreq[int64(el[1])]--
				mapFreqOfFreq[mapFreq[int64(el[1])]]++

				if val, found := mapFreqOfFreq[mapFreq[int64(el[1])]+1]; found && val > 0 {
					mapFreqOfFreq[mapFreq[int64(el[1])]+1]--
				}
			}
		} else if el[0] == 3 {
			if mapFreqOfFreq[int64(el[1])] > 0 {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}
		}
	}

	return result
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
