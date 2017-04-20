package main

// https://www.hackerrank.com/challenges/marcs-cakewalk
// 10 min

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	ar := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&ar[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ar)))

	result := float64(0)
	for i := 0; i < N; i++ {
		result += float64(ar[i]) * math.Pow(2, float64(i))
	}

	fmt.Printf("%0.f", result)

}
