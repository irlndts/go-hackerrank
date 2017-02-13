package main

import (
	"fmt"
	"sort"
)

// https://www.hackerrank.com/challenges/mini-max-sum

/*
	func main() {
		    fmt.Println(MiniMaxSum())
		}
*/

func MiniMaxSum() (int, int) {
	ar := make([]int, 5)
	for i := 0; i < len(ar); i++ {
		fmt.Scan(&ar[i])
	}

	sort.Ints(ar)

	min := 0
	for i := 0; i < len(ar)-1; i++ {
		min += ar[i]
	}

	max := 0
	for i := 1; i < len(ar); i++ {
		max += ar[i]
	}

	return min, max
}
