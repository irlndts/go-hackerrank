package main

// https://www.hackerrank.com/contests/w21/challenges/luck-balance

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	var result int
	fmt.Scanf("%v %v", &N, &K)
	var ar []int
	for i := 0; i < N; i++ {
		var L, T int
		fmt.Scanf("%v %v", &L, &T)
		if T == 1 {
			ar = append(ar, L)
		} else {
			result += L
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ar)))
	for i := 0; i < K; i++ {
		result += ar[i]
	}
	for i := K; i < len(ar); i++ {
		result -= ar[i]
	}
	fmt.Println(result)
}
