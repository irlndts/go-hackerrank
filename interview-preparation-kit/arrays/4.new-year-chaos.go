package main

import "fmt"

// https://www.hackerrank.com/challenges/new-year-chaos/problem

func minimumBribes(q []int32) {
	result := int32(0)
	for i := int32(len(q)) - 1; i >= 0; i-- {
		if q[i]-(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}
		for j := max(0, q[i]-2); j < i; j++ {
			if q[j] > q[i] {
				result++
			}
		}
	}
	fmt.Println(result)
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
