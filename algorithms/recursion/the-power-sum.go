package main

// https://www.hackerrank.com/challenges/the-power-sum
// 20 mins

import (
	"fmt"
	"math"
)

var (
	result = 0
	N      = float64(0)
	X      = float64(0)
)

func main() {
	fmt.Scan(&X)
	fmt.Scan(&N)
	for i := float64(1); i <= math.Sqrt(X); i++ {
		solve(i, 0)
	}
	fmt.Println(result)
}

func solve(i, c float64) {
	tmp := math.Pow(i, N)
	current := tmp + c

	if current == X {
		result++
	} else if current < X {
		for j := i + 1; j <= math.Sqrt(X); j++ {
			solve(j, current)
		}
	}
}
