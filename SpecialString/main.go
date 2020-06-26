package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "asasd" //expected result 7
	s := "aabaa"
	fmt.Println(substrCount(int32(len(s)), s))
}

func substrCount(n int32, s string) int64 {
	count := int64(0)
	stringArr := strings.Split(s, "")

	for i := 0; i < len(stringArr); i++ {
		startChar := stringArr[i]
		diffCharIdx := -1

		for j := i + 1; j < len(stringArr); j++ {
			currChar := stringArr[j]
			if startChar == currChar {
				if (diffCharIdx == -1) || (j-diffCharIdx) == (diffCharIdx-i) {
					count++
				}
			} else {
				if diffCharIdx == -1 {
					diffCharIdx = j
				} else {
					break
				}
			}
		}
	}

	return int64(len(stringArr)) + count
}

// static long substrCount(string s)
// {
// 	long retVal = s.Length;

// 	for (int i = 0; i < s.Length; i++)
// 	{
// 		var startChar = s[i];
// 		int diffCharIdx = -1;
// 		for (int j = i + 1; j < s.Length; j++)
// 		{
// 			var currChar = s[j];
// 			if (startChar == currChar)
// 			{
// 				if ((diffCharIdx == -1) ||
// 					(j - diffCharIdx) == (diffCharIdx - i))
// 					retVal++;
// 			}
// 			else
// 			{
// 				if (diffCharIdx == -1)
// 					diffCharIdx = j;
// 				else
// 					break;
// 			}
// 		}
// 	}
// 	return retVal;
// }
