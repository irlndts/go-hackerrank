package main

import "sort"

// https://www.hackerrank.com/challenges/equality-in-a-array/problem

func equalizeArray(arr []int32) int32 {
	m := make(map[int]int)
	for _, s := range arr {
		m[int(s)]++
	}

	values := make([]int, 0, len(m))
	for _, s := range m {
		values = append(values, s)
	}

	result := 0
	sort.Ints(values)
	for i := 0; i < len(values)-1; i++ {
		result += values[i]
	}
	return int32(result)
}
