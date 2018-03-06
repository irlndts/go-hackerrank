package main

// https://www.hackerrank.com/contests/101hack53/challenges/cut-board
// 2D Array
// Covering

import (
	"fmt"
)

func CutBoard(n, m, x, y int) {
	if (n*m-x-y)%2 == 1 {
		fmt.Println("NO")
		return
	}

	var ar [][]int
	ar = make([][]int, n)
	for i := 0; i < n; i++ {
		ar[i] = make([]int, m)
	}

	for i := 0; i < x; i++ {
		ar[0][i] = 1
	}

	for i := m - y; i < m; i++ {
		ar[n-1][i] = 1
	}

	for i := m - y; i < m; i++ {
		ar[n-1][i] = 1
	}

	var result []string
	result = append(result, "YES")
	result = append(result, fmt.Sprint((n*m-x-y)/2))
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ar[i][j] == 1 {
				continue
			}

			if (j != m-1 && ar[i][j+1] != 1) &&
				!(i == n-2 && ar[i+1][j] != 1) {
				ar[i][j] = 1
				ar[i][j+1] = 1
				result = append(result, fmt.Sprint(i+1, j+1, i+1, j+2))
			} else if ar[i+1][j] != 1 {
				ar[i][j] = 1
				ar[i+1][j] = 1
				result = append(result, fmt.Sprint(i+1, j+1, i+2, j+1))
			} else {
				fmt.Println("NO")
				return
			}
		}
	}

	for _, r := range result {
		fmt.Println(r)
	}
}
