package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s1 := "ABCDEF"
	s2 := "FBDAMN"

	fmt.Println(commonChild(s1, s2))
}

func commonChild(s1 string, s2 string) int32 {
	return int32(lcs(strings.Split(s1, ""), strings.Split(s2, ""), len(s1), len(s2)))
}

func lcs(s1Arr []string, s2Arr []string, s1ArrLen int, s2ArrLen int) int64 {
	L := make([][]int64, s1ArrLen+1)
	for i := 0; i < len(L); i++ {
		L[i] = make([]int64, s2ArrLen+1)
	}

	for i := 0; i <= s1ArrLen; i++ {
		for j := 0; j <= s2ArrLen; j++ {
			if i == 0 || j == 0 {
				L[i][j] = 0
			} else if s1Arr[i-1] == s2Arr[j-1] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				L[i][j] = int64(math.Max(float64(L[i-1][j]), float64(L[i][j-1])))
			}
		}
	}

	return L[s1ArrLen][s2ArrLen]
}

// /* Returns length of LCS for X[0..m-1], Y[0..n-1] */
// int lcs( char *X, char *Y, int m, int n )
// {
//    int L[m+1][n+1];
//    int i, j;

//    /* Following steps build L[m+1][n+1] in bottom up fashion. Note
//       that L[i][j] contains length of LCS of X[0..i-1] and Y[0..j-1] */
//    for (i=0; i<=m; i++)
//    {
//      for (j=0; j<=n; j++)
//      {
//        if (i == 0 || j == 0)
//          L[i][j] = 0;

//        else if (X[i-1] == Y[j-1])
//          L[i][j] = L[i-1][j-1] + 1;

//        else
//          L[i][j] = max(L[i-1][j], L[i][j-1]);
//      }
//    }

//    /* L[m][n] contains length of LCS for X[0..n-1] and Y[0..m-1] */
//    return L[m][n];
// }
