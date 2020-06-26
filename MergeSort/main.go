package main

import "fmt"

func main() {
	arr := []int32{1, 5, 3, 7} // expected result 1
	//arr := []int32{7, 5, 3, 1} // expected result 6

	fmt.Println(countInversions(arr))
}

func countInversions(arr []int32) int64 {
	countInv := int64(0)

	mergeSort(arr, &countInv)

	return countInv
}

func mergeSort(arr []int32, count *int64) []int32 {
	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle], count)
	right := mergeSort(arr[middle:], count)

	return merge(left, right, count)
}

func merge(left []int32, right []int32, count *int64) []int32 {
	temp := []int32{}
	for i := 0; i < len(left)+len(right); i++ {
		temp = append(temp, 0)
	}

	leftIndex, rightIndex, index := 0, 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			temp[index] = left[leftIndex]
			index++
			leftIndex++
		} else {
			temp[index] = right[rightIndex]
			*count += int64((rightIndex + len(left)) - index)
			index++
			rightIndex++
		}
	}

	for j := leftIndex; j < len(left); j++ {
		temp[index] = left[j]
		index++
	}

	for k := rightIndex; k < len(right); k++ {
		temp[index] = right[k]
		index++
	}

	return temp
}
